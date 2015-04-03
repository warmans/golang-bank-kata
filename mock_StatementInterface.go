package main

import "github.com/stretchr/testify/mock"

type MockStatementInterface struct {
	mock.Mock
}

func (m *MockStatementInterface) Print(transactions []Transaction) {
	m.Called(transactions)
}
