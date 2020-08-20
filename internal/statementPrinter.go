package internal

import (
	"log"
	"math"
)

type StatementLine struct {
	transaction Transaction
	currentBalance float32
}

type StatementPrinter struct{
	logger *log.Logger
}

func NewStatementLine(t Transaction, cb float32) StatementLine {
	return StatementLine{transaction:t, currentBalance:cb}
}

func NewStatementPrinter(loggerInjected *log.Logger) *StatementPrinter {
	return &StatementPrinter{logger:loggerInjected}
}

func (s *StatementPrinter) PrintAllStatements(statementLines [] StatementLine) {
	s.printStatementHeaders()
	s.printArrayOfStatements(statementLines)
}

func (s *StatementPrinter) printStatementHeaders() {
	statementHeaders := "date || credit || debit || balance"
	s.logger.Print(statementHeaders)
}

func (s *StatementPrinter) printArrayOfStatements(statementLines []StatementLine) {
	for i := len(statementLines) - 1; i >= 0; i-- {
		statementLine := statementLines[i]
		if statementLine.transaction.Amount > 0 {
			s.logger.Printf("%s || %.0f || || %.0f\n", statementLine.transaction.Date,
				statementLine.transaction.Amount, statementLine.currentBalance)
		} else {
			s.logger.Printf("%s || || %.0f || %.0f\n", statementLine.transaction.Date,
				math.Abs(float64(statementLine.transaction.Amount)), statementLine.currentBalance)
		}
	}
}