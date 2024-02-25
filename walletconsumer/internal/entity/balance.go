package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID        string
	AccountID string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBalance(accountId string, amount float64) (*Balance, error) {
	balance := &Balance{
		ID:        uuid.New().String(),
		AccountID: accountId,
		Amount:    amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := balance.Validate()
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (b *Balance) UpdateBalance(accountId string, amount float64) error {
	b.AccountID = accountId
	b.Amount = amount
	b.UpdatedAt = time.Now()
	err := b.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (b *Balance) Validate() error {
	if b.AccountID == "" {
		return errors.New("Account ID is required")
	}
	if b.Amount <= 0.0 {
		return errors.New("Amount must be greater than zero")
	}
	return nil
}
