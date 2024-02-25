package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterNewBalance(t *testing.T) {
	balance, err := NewBalance("023b19f6-c81f-497a-bdc2-4602e7856632", 1000.00)
	assert.Nil(t, err)
	assert.NotNil(t, balance)
	assert.NotNil(t, balance.ID)
	assert.Equal(t, "023b19f6-c81f-497a-bdc2-4602e7856632", balance.AccountID)
	assert.Equal(t, 1000.00, balance.Amount)
}

func TestRegisterNewBalanceWhenArgsAreInvalid(t *testing.T) {
	balance, err := NewBalance("", -10.00)
	assert.NotNil(t, err)
	assert.Nil(t, balance)
}

func TestUpdateBalance(t *testing.T) {
	balance, _ := NewBalance("023b19f6-c81f-497a-bdc2-4602e7856632", 1000.00)
	err := balance.UpdateBalance("023b19f6-c81f-497a-bdc2-4602e7856632", 2000.00)
	assert.Nil(t, err)
	assert.Equal(t, "023b19f6-c81f-497a-bdc2-4602e7856632", balance.AccountID)
	assert.Equal(t, 2000.00, balance.Amount)
}

func TestUpdateBalanceWhenArgsAreInvalid(t *testing.T) {
	balance, _ := NewBalance("023b19f6-c81f-497a-bdc2-4602e7856632", 1000.00)
	err := balance.UpdateBalance("023b19f6-c81f-497a-bdc2-4602e7856632", -5.0)
	assert.Error(t, err, "Amount must be greater than zero")
}
