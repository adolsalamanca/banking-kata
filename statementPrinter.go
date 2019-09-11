package banking_kata

import "log"

type StatementPrinter struct{
	logger *log.Logger
}

func NewStatementPrinter(loggerInjected *log.Logger) *StatementPrinter {
	return &StatementPrinter{logger:loggerInjected}
}

func (s *StatementPrinter) PrintStatement(){
	s.logger.Println("Statement")
}