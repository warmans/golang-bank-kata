package main
import "time"

type Transaction struct {
    Date   time.Time
    Amount int
}

type TransactionStoreInterface interface {
    StoreDeposit(amount int)
    StoreWithdrawal(amount int)
    GetTransactions() []Transaction
}

type TransactionStore struct {
    clock        ClockInterface
    transactions []Transaction
}

func (ts *TransactionStore) StoreDeposit(amount int) {
    ts.transactions = append(ts.transactions, Transaction{Date: ts.clock.GetNow(), Amount: amount})
}

func (ts *TransactionStore) StoreWithdrawal(amount int) {
    ts.transactions = append(ts.transactions, Transaction{Date: ts.clock.GetNow(), Amount: (0-amount)})
}

func (ts *TransactionStore) GetTransactions() []Transaction {
    return ts.transactions
}
