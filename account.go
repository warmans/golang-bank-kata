package main

type AccountInterface interface {
    Deposit(amount int)
    Withdraw(amount int)
    PrintStatement()
}

type Account struct {
    TransactionStore TransactionStoreInterface
    Statement        StatementInterface
}

func (acc *Account) Deposit(amount int) {
    acc.TransactionStore.StoreDeposit(amount)
}

func (acc *Account) Withdraw(amount int) {
    acc.TransactionStore.StoreWithdrawal(amount)
}

func (acc *Account) PrintStatement() {
    acc.Statement.Print(acc.TransactionStore.GetTransactions())
}
