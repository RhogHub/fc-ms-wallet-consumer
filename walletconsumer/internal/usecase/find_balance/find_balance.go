package find_balance

import (
	"time"

	"github.com/RhogHub/fc-ms-wallet-consumer/internal/gateway"
)

type FindBalanceInputDTO struct {
	AccountID string `json:"account_id"`
}

type FindBalanceOutputDTO struct {
	ID        string
	AccountID string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FindBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewFindBalanceUseCase(balanceGateway gateway.BalanceGateway) *FindBalanceUseCase {
	return &FindBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (uc *FindBalanceUseCase) Execute(input FindBalanceInputDTO) (*FindBalanceOutputDTO, error) {
	balance, err := uc.BalanceGateway.FindBalanceByAccountID(input.AccountID)
	if err != nil {
		return nil, err
	}

	output := &FindBalanceOutputDTO{
		ID:        balance.ID,
		AccountID: balance.AccountID,
		Amount:    balance.Amount,
		CreatedAt: balance.CreatedAt,
		UpdatedAt: balance.UpdatedAt,
	}

	return output, nil
}
