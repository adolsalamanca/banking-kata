package banking_kata

import (
	"log"
	"math"
)

type StatementPrinter struct{
	logger *log.Logger
}

type StatementLine struct {
	transaction Transaction
	currentBalance float32
}

func NewStatementLine(t Transaction, cb float32) StatementLine {
	return StatementLine{transaction:t, currentBalance:cb}
}

func NewStatementPrinter(loggerInjected *log.Logger) *StatementPrinter {
	return &StatementPrinter{logger:loggerInjected}
}

func (s *StatementPrinter) PrintStatementHeaders() {
	statementHeaders := "date || credit || debit || balance"
	s.logger.Print(statementHeaders)
}

func (s *StatementPrinter) PrintStatement(lines [] StatementLine) {

	s.PrintStatementHeaders()
	for i:= len(lines)-1; i>=0; i-- {
		statementLine := lines[i]
		if statementLine.transaction.Amount > 0 {
			s.logger.Printf("%s || %.0f || || %.0f\n", statementLine.transaction.Date,
				statementLine.transaction.Amount, statementLine.currentBalance)
		} else {
			s.logger.Printf("%s || || %.0f || %.0f\n", statementLine.transaction.Date,
				math.Abs(float64(statementLine.transaction.Amount)), statementLine.currentBalance)
		}
	}

}

