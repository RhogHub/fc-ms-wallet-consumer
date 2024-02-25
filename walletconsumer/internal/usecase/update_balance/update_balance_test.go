package update_balance

import (
	"testing"

	"github.com/RhogHub/fc-ms-wallet-consumer/internal/entity"
	"github.com/RhogHub/fc-ms-wallet-consumer/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateBalanceUseCase_Execute(t *testing.T) {
	balance, _ := entity.NewBalance("023b19f6-c81f-497a-bdc2-4602e7856632", 1000.00)
	balanceMock := &mocks.BalanceGatewayMock{}
	balanceMock.On("NewBalance").Return(nil)

	m := &mocks.BalanceGatewayMock{}
	m.On("UpdateBalance", mock.Anything).Return(nil)
	uc := NewUpdateUseCase(m)

	output, err := uc.Execute(UpdateBalanceInputDTO{
		AccountID: balance.AccountID,
		Amount:    2000.00,
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "023b19f6-c81f-497a-bdc2-4602e7856632", output.AccountID)
	assert.Equal(t, 2000.00, output.Amount)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "UpdateBalance", 1)
}
