package update_balance

import (
	"time"

	"github.com/RhogHub/fc-ms-wallet-consumer/internal/entity"
	"github.com/RhogHub/fc-ms-wallet-consumer/internal/gateway"
)

type UpdateBalanceInputDTO struct {
	AccountID string
	Amount    float64
}

type UpdateBalanceOutputDTO struct {
	ID        string
	AccountID string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewUpdateUseCase(balanceGateway gateway.BalanceGateway) *UpdateBalanceUseCase {
	return &UpdateBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (uc *UpdateBalanceUseCase) Execute(input UpdateBalanceInputDTO) (*UpdateBalanceOutputDTO, error) {
	balance := &entity.Balance{}
	err := balance.UpdateBalance(input.AccountID, input.Amount)
	if err != nil {
		return nil, err
	}
	err = uc.BalanceGateway.UpdateBalance(balance)
	if err != nil {
		return nil, err
	}

	output := &UpdateBalanceOutputDTO{
		ID:        balance.ID,
		AccountID: balance.AccountID,
		Amount:    balance.Amount,
		CreatedAt: balance.CreatedAt,
		UpdatedAt: balance.UpdatedAt,
	}
	return output, nil
}
