package temp

// import (
// 	"fmt"

// 	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
// )

// func main() {
// 	msgCHan := make(chan *ckafka.Message)
// 	topics := []string{"balances"}
// 	servers := "kafka:29092"

// 	go Consume(topics, servers, msgCHan) // faz a mágica.

// 	for {
// 		msg := <-msgCHan
// 		println(string(msg.Value))
// 	}
// }

// func Consume(topics []string, servers string, msgChan chan *ckafka.Message) {
// 	kafkaConsumer, err := ckafka.NewConsumer(&ckafka.ConfigMap{
// 		"bootstrap.servers": servers,
// 		"group.id":          "wallet",
// 		"auto.offset.reset": "earliest",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = kafkaConsumer.SubscribeTopics(topics, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for {
// 		msg, err := kafkaConsumer.ReadMessage(-1)
// 		if err == nil {
// 			msgChan <- msg
// 			fmt.Printf("Mensagem recebida: %s\n", string(msg.Value))
// 		}
// 	}
// }
