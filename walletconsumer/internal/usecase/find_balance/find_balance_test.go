package find_balance

import (
	"testing"

	"github.com/RhogHub/fc-ms-wallet-consumer/internal/entity"
	"github.com/RhogHub/fc-ms-wallet-consumer/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindBalanceUseCase_Execute(t *testing.T) {
	balance, _ := entity.NewBalance("023b19f6-c81f-497a-bdc2-4602e7856632", 1000)
	balanceMock := &mocks.BalanceGatewayMock{}
	balanceMock.On("NewBalance").Return(nil)

	m := &mocks.BalanceGatewayMock{}
	m.On("FindBalanceByAccountID", mock.Anything).Return(balance, nil)
	uc := NewFindBalanceUseCase(m)

	inputDto := FindBalanceInputDTO{
		AccountID: balance.AccountID,
	}

	output, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "023b19f6-c81f-497a-bdc2-4602e7856632", output.AccountID)
	assert.Equal(t, 1000.00, output.Amount)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindBalanceByAccountID", 1)
}
