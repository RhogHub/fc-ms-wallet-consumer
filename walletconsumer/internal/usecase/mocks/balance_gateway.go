package mocks

import (
	"github.com/RhogHub/fc-ms-wallet-consumer/internal/entity"
	"github.com/stretchr/testify/mock"
)

type BalanceGatewayMock struct {
	mock.Mock
}

func (m *BalanceGatewayMock) NewBalance(balance *entity.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

func (m *BalanceGatewayMock) FindBalanceByAccountID(accountId string) (*entity.Balance, error) {
	args := m.Called(accountId)
	return args.Get(0).(*entity.Balance), args.Error(1)
}

func (m *BalanceGatewayMock) UpdateBalance(balance *entity.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}
