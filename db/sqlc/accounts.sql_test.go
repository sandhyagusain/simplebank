package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/techschool/simplebank/db/util"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, account)

	assert.Equal(t, arg.Owner, account.Owner)
	assert.Equal(t, arg.Currency, account.Currency)
	assert.Equal(t, arg.Balance, account.Balance)

	assert.NotZero(t, account.ID)
	assert.NotZero(t, account.CreatedAt)

	return account
}

func TestQueries_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestQueries_DeleteAccount(t *testing.T) {
	err := testQueries.DeleteAccount(context.Background(), 78)
	assert.NoError(t, err)
}

func TestQueries_GetAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, account2)

	assert.Equal(t, account1.ID, account2.ID)
	assert.Equal(t, account1.Owner, account2.Owner)
	assert.Equal(t, account1.Currency, account2.Currency)
	assert.Equal(t, account1.Balance, account2.Balance)
}

func TestQueries_UpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, account2)

	assert.Equal(t, account1.ID, account2.ID)
	assert.Equal(t, account1.Owner, account2.Owner)
	assert.Equal(t, account1.Currency, account2.Currency)
	assert.Equal(t, arg.Balance, account2.Balance)
}
