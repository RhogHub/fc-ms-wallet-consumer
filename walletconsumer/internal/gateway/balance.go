package gateway

import "github.com/RhogHub/fc-ms-wallet-consumer/internal/entity"

type BalanceGateway interface {
	NewBalance(balance *entity.Balance) error
	UpdateBalance(balance *entity.Balance) error
	FindBalanceByAccountID(account_id string) (*entity.Balance, error)
}
