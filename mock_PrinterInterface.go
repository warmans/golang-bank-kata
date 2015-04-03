package main

import "github.com/stretchr/testify/mock"

type MockPrinterInterface struct {
	mock.Mock
}

func (m *MockPrinterInterface) Printf(format string, values ...interface{}) {
	m.Called(format, values)
}
