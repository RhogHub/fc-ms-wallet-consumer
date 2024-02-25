package database

import (
	"database/sql"

	"github.com/RhogHub/fc-ms-wallet-consumer/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (b *BalanceDB) FindBalanceByAccountID(account_id string) (*entity.Balance, error) {
	balance := &entity.Balance{}
	stmt, err := b.DB.Prepare("SELECT id, account_id, balance FROM balances WHERE account_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(account_id)
	if err := row.Scan(&balance.ID, &balance.AccountID, &balance.Amount); err != nil {
		if err == sql.ErrNoRows {
			// A conta não existe
			return nil, nil
		}
		return nil, err
	}
	return balance, nil
}

func (b *BalanceDB) NewBalance(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("INSERT INTO balances (id, account_id, balance, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&balance.ID, &balance.AccountID, &balance.Amount, &balance.CreatedAt, &balance.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (b *BalanceDB) UpdateBalance(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("UPDATE balances SET balance=?, updated_at=? WHERE account_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&balance.Amount, &balance.UpdatedAt, &balance.AccountID)
	if err != nil {
		return err
	}
	return nil
}
