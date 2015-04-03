package main

import "fmt"

type StatementInterface interface {
    Print(transactions []Transaction)
}

type PrinterInterface interface {
    Printf(format string, values... interface{})
}

type StdOutPrinter struct {}

func (p *StdOutPrinter) Printf(format string, values... interface{}) {
    fmt.Printf(format, values...)
}

type BufferPrinter struct {
    buffer string
}

func (p *BufferPrinter) Printf(format string, values... interface{}) {
    p.buffer += fmt.Sprintf(format, values...)
}

func (p *BufferPrinter) GetBuffer() string {
    return p.buffer
}

type Statement struct {
    printer PrinterInterface
}

func (s *Statement) Print(transactions []Transaction) {
    s.printer.Printf("Date | Amount | Balance\n")
    balance := 0
    for _, transaction := range transactions {
        balance += transaction.Amount
        s.printer.Printf("%s | %d | %d\n", transaction.Date.Format("2006-01-02"), transaction.Amount, balance)
    }
}
