package main
import (
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func TestAppendsDateToDeposit(t *testing.T) {
    now := time.Now()

    clock := new(MockClockInterface)
    clock.On("GetNow").Return(now)

    store := TransactionStore{clock: clock}
    store.StoreDeposit(100)

    assert.Equal(t, now, store.GetTransactions()[0].Date)
}

func TestAppendsDateToWithdrawal(t *testing.T) {
    now := time.Now()

    clock := new(MockClockInterface)
    clock.On("GetNow").Return(now)

    store := TransactionStore{clock: clock}
    store.StoreWithdrawal(100)

    assert.Equal(t, now, store.GetTransactions()[0].Date)
}

func TestStoresGivenDepositAmountAsPositive(t *testing.T) {
    store := TransactionStore{clock: &Clock{}}
    store.StoreDeposit(100)

    assert.Equal(t, 100, store.GetTransactions()[0].Amount)
}

func TestStoresGivenWithdrawalAmountAsNegative(t *testing.T) {
    store := TransactionStore{clock: &Clock{}}
    store.StoreWithdrawal(100)

    assert.Equal(t, -100, store.GetTransactions()[0].Amount)
}

func TestStoresMultipleTransactions(t *testing.T) {
    store := TransactionStore{clock: &Clock{}}
    store.StoreDeposit(100)
    store.StoreWithdrawal(200)
    store.StoreDeposit(300)

    transactions := store.GetTransactions()

    assert.Len(t, transactions, 3)
}
