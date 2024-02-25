package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/RhogHub/fc-ms-wallet-consumer/internal/database"
	"github.com/RhogHub/fc-ms-wallet-consumer/internal/usecase/create_balance"
	"github.com/RhogHub/fc-ms-wallet-consumer/internal/usecase/find_balance"
	"github.com/RhogHub/fc-ms-wallet-consumer/internal/usecase/update_balance"
	"github.com/RhogHub/fc-ms-wallet-consumer/internal/web"
	"github.com/RhogHub/fc-ms-wallet-consumer/internal/web/webserver"
	"github.com/RhogHub/fc-ms-wallet-consumer/pkg/kafka"
	"github.com/RhogHub/fc-ms-wallet-consumer/pkg/uow"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func CheckDBConnection(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}
	return nil
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "db_consumer"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	balanceDb := database.NewBalanceDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("BalanceDB", func(tx *sql.Tx) interface{} {
		return database.NewBalanceDB(db)
	})

	createBalanceUseCase := create_balance.NewCreateBalanceUseCase(balanceDb)
	findBalanceUseCase := find_balance.NewFindBalanceUseCase(balanceDb)
	updateBalanceUseCase := update_balance.NewUpdateUseCase(balanceDb)

	webserver := webserver.NewWebServer(":3003")

	balanceHandler := web.NewWebFindBalanceHandler(*findBalanceUseCase)

	webserver.AddHandler("/balances/{account_id}", balanceHandler.FindBalance)

	if err := CheckDBConnection(db); err != nil {
		panic(err)
	}

	fmt.Println("Server is running")
	webserver.Start()

	logger := log.New(os.Stdout, "KafkaConsumer: ", log.LstdFlags)
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	msgChan := make(chan *ckafka.Message)
	topics := []string{"balances"}

	kafkaConsumer := kafka.NewConsumer(&configMap, topics)

	go kafkaConsumer.Consume(msgChan)

	for {
		msg := <-msgChan

		logger.Printf("Mensagem recebida: %s\n", string(msg.Value))
		type KafkaMessage struct {
			Name    string `json:"Name"`
			Payload struct {
				AccountIDFrom        string  `json:"account_id_from"`
				AccountIDTo          string  `json:"account_id_to"`
				BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
				BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
			} `json:"Payload"`
		}

		var kafkaMsg *KafkaMessage
		err := json.Unmarshal(msg.Value, &kafkaMsg)
		if err != nil {
			logger.Printf("Erro ao decodificar mensagem JSON: %v\n", err)
			return
		}

		if kafkaMsg.Payload.AccountIDTo == kafkaMsg.Payload.AccountIDFrom {
			logger.Printf("AccountIDTo = AccountIDFrom")
			kafkaMsg.Payload.BalanceAccountIDTo = kafkaMsg.Payload.BalanceAccountIDFrom
		}

		checkIfAccountFromExist, _ := findBalanceUseCase.BalanceGateway.FindBalanceByAccountID(kafkaMsg.Payload.AccountIDFrom)
		if checkIfAccountFromExist == nil {
			logger.Printf("Create accountFrom")
			createBalanceUseCase.Execute(create_balance.CreateBalanceInputDTO{
				AccountID: kafkaMsg.Payload.AccountIDFrom,
				Amount:    kafkaMsg.Payload.BalanceAccountIDFrom,
			})

		} else {
			logger.Printf("Update accountFrom")
			updateBalanceUseCase.Execute(update_balance.UpdateBalanceInputDTO{
				AccountID: kafkaMsg.Payload.AccountIDFrom,
				Amount:    kafkaMsg.Payload.BalanceAccountIDFrom,
			})
		}
		if err != nil {
			panic(err)
		}

		checkIfAccountToExist, _ := findBalanceUseCase.BalanceGateway.FindBalanceByAccountID(kafkaMsg.Payload.AccountIDTo)
		if checkIfAccountToExist == nil {
			logger.Printf("Create accountTo")
			createBalanceUseCase.Execute(create_balance.CreateBalanceInputDTO{
				AccountID: kafkaMsg.Payload.AccountIDTo,
				Amount:    kafkaMsg.Payload.BalanceAccountIDTo,
			})
		} else {
			logger.Printf("Update accountTo")
			updateBalanceUseCase.Execute(update_balance.UpdateBalanceInputDTO{
				AccountID: kafkaMsg.Payload.AccountIDTo,
				Amount:    kafkaMsg.Payload.BalanceAccountIDTo,
			})
		}
		if err != nil {
			panic(err)
		}
	}
}
