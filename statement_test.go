package main
import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
)

func TestPrintsTitleRow(t *testing.T) {
    printer := &BufferPrinter{}

    statement := Statement{printer: printer}
    statement.Print(make([]Transaction, 0))

    assert.Equal(t, printer.GetBuffer(), "Date | Amount | Balance\n")
}

func TestPrintsTransactionRows(t *testing.T) {

    date := time.Now()
    printer := &BufferPrinter{}

    transactions := make([]Transaction, 2)
    transactions[0] = Transaction{Date: date, Amount: 200}
    transactions[1] = Transaction{Date: date, Amount: -100}

    statement := Statement{printer: printer}
    statement.Print(transactions)

    assert.Contains(t, printer.GetBuffer(), date.Format("2006-01-02") + " | 200 | 200\n")
    assert.Contains(t, printer.GetBuffer(), date.Format("2006-01-02") + " | -100 | 100\n")
}
