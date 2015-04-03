package main
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "time"
)

func TestAcceptanceTransactionPrinting(t *testing.T) {

    mockClock := new(MockClockInterface)
    mockClock.On("GetNow").Return(time.Parse(time.RFC3339, "2015-01-01T00:00:00+00:00"))

    printer := &BufferPrinter{}
    statement := &Statement{printer: printer}
    transactionStore := &TransactionStore{clock: mockClock}
    account := Account{TransactionStore: transactionStore, Statement: statement}

    account.Deposit(200)
    account.Deposit(200)
    account.Withdraw(100)
    account.PrintStatement()

    assert.Equal(t, `Date | Amount | Balance
2015-01-01 | 200 | 200
2015-01-01 | 200 | 400
2015-01-01 | -100 | 300
`, printer.GetBuffer())

}
