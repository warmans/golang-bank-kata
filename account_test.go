package main

import (
    "testing"
)

func TestDepositStoresTransaction(t *testing.T) {

    transactionStore := new(MockTransactionStoreInterface)
    transactionStore.On("StoreDeposit", 100).Return()

    acc := Account{TransactionStore: transactionStore}
    acc.Deposit(100)

    transactionStore.AssertCalled(t, "StoreDeposit", 100)
}

func TestWithdrawStoresTransaction(t *testing.T) {

    transactionStore := new(MockTransactionStoreInterface)
    transactionStore.On("StoreWithdrawal", 100).Return()

    acc := Account{TransactionStore: transactionStore}
    acc.Withdraw(100)

    transactionStore.AssertCalled(t, "StoreWithdrawal", 100)
}

func TestPrintStatementInvokesPrinter(t *testing.T) {

    statement := new(MockStatementInterface)
    statement.On("Print", make([]Transaction, 0)).Return()

    transactionStore := &TransactionStore{transactions: make([]Transaction, 0)}

    acc := Account{TransactionStore: transactionStore, Statement: statement}
    acc.PrintStatement()

    statement.AssertCalled(t, "Print", make([]Transaction, 0))
}
