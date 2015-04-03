package main

import "github.com/stretchr/testify/mock"

import "time"

type MockClockInterface struct {
	mock.Mock
}

func (m *MockClockInterface) GetNow() time.Time {
	ret := m.Called()

	r0 := ret.Get(0).(time.Time)

	return r0
}
