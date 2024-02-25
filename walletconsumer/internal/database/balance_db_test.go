package database

import (
	"database/sql"
	"testing"

	"github.com/RhogHub/fc-ms-wallet-consumer/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type BalanceDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	BalanceDB *BalanceDB
	balance   *entity.Balance
}

func (s *BalanceDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table balances (id varchar(255), account_id varchar(255), balance int, created_at date, updated_at date)")
	s.BalanceDB = NewBalanceDB(db)
	s.balance, _ = entity.NewBalance("023b19f6-c81f-497a-bdc2-4602e7856632", 1000.00)
}

func (s *BalanceDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE balances")
}

func TestBalanceDBTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDBTestSuite))
}

func (s *BalanceDBTestSuite) TestNewBalance() {
	balance := &entity.Balance{
		ID:        "da5151aa-8c22-455c-aae0-9150efd12b3b",
		AccountID: "023b19f6-c81f-497a-bdc2-4602e7856632",
		Amount:    666.66,
	}
	err := s.BalanceDB.NewBalance(balance)
	s.Nil(err)
}

func (s *BalanceDBTestSuite) TestFindBalanceByAccountID() {
	balance, _ := entity.NewBalance("023b19f6-c81f-497a-bdc2-4602e7856632", 1000.00)
	s.BalanceDB.NewBalance(balance)

	balanceDB, err := s.BalanceDB.FindBalanceByAccountID(balance.AccountID)
	s.Nil(err)
	s.Equal(balance.ID, balanceDB.ID)
	s.Equal(balance.AccountID, balanceDB.AccountID)
	s.Equal(balance.Amount, balanceDB.Amount)
}

func (s *BalanceDBTestSuite) TestUpdateBalance() {
	balance, _ := entity.NewBalance(s.balance.AccountID, s.balance.Amount)
	err := s.BalanceDB.NewBalance(balance)
	s.Nil(err)

	err = balance.UpdateBalance("023b19f6-c81f-497a-bdc2-4602e7856632", 4000.00)
	if err != nil {
		panic(err)
	}

	err = s.BalanceDB.UpdateBalance(balance)
	s.Nil(err)

	balanceDB, err := s.BalanceDB.FindBalanceByAccountID(balance.AccountID)
	s.Nil(err)
	//s.Equal(balance.ID, balanceDB.ID)
	s.Equal(balance.AccountID, balanceDB.AccountID)
	s.Equal(balance.Amount, balanceDB.Amount)
	s.Equal(balance.Amount, 4000.00)
	s.Equal(balanceDB.Amount, 4000.00)
	s.NotEqual(balanceDB.Amount, s.balance.Amount)
}
