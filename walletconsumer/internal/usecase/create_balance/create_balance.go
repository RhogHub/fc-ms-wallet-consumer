package create_balance

import (
	"time"

	"github.com/RhogHub/fc-ms-wallet-consumer/internal/entity"
	"github.com/RhogHub/fc-ms-wallet-consumer/internal/gateway"
)

type CreateBalanceInputDTO struct {
	AccountID string
	Amount    float64
}

type CreateBalanceOutputDTO struct {
	ID        string
	AccountID string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewCreateBalanceUseCase(balanceGateway gateway.BalanceGateway) *CreateBalanceUseCase {
	return &CreateBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (uc *CreateBalanceUseCase) Execute(input CreateBalanceInputDTO) (*CreateBalanceOutputDTO, error) {
	balance, err := entity.NewBalance(input.AccountID, input.Amount)
	if err != nil {
		return nil, err
	}
	err = uc.BalanceGateway.NewBalance(balance)
	if err != nil {
		return nil, err
	}

	output := &CreateBalanceOutputDTO{
		ID:        balance.ID,
		AccountID: balance.AccountID,
		Amount:    balance.Amount,
		CreatedAt: balance.CreatedAt,
		UpdatedAt: balance.UpdatedAt,
	}
	return output, nil
}
