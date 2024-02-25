package create_balance

import (
	"testing"

	"github.com/RhogHub/fc-ms-wallet-consumer/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBalanceUseCase_Execute(t *testing.T) {
	m := &mocks.BalanceGatewayMock{}
	m.On("NewBalance", mock.Anything).Return(nil)
	uc := NewCreateBalanceUseCase(m)

	output, err := uc.Execute(CreateBalanceInputDTO{
		AccountID: "023b19f6-c81f-497a-bdc2-4602e7856632",
		Amount:    1000.00,
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "023b19f6-c81f-497a-bdc2-4602e7856632", output.AccountID)
	assert.Equal(t, 1000.00, output.Amount)
	m.AssertExpectations(t) //Garante q o NewBalance foi chamado
	m.AssertNumberOfCalls(t, "NewBalance", 1)
}
