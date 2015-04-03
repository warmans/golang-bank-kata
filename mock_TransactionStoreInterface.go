package main

import "github.com/stretchr/testify/mock"

type MockTransactionStoreInterface struct {
	mock.Mock
}

func (m *MockTransactionStoreInterface) StoreDeposit(amount int) {
	m.Called(amount)
}
func (m *MockTransactionStoreInterface) StoreWithdrawal(amount int) {
	m.Called(amount)
}
func (m *MockTransactionStoreInterface) GetTransactions() []Transaction {
	ret := m.Called()

	var r0 []Transaction
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]Transaction)
	}

	return r0
}
