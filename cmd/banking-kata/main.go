package main

import (
	"github.com/adolsalamanca/banking-kata/internal"
	"log"
	"os"
)

func main() {
	logger := &log.Logger{}
	logger.SetOutput(os.Stdout)

	account := internal.NewAccount(internal.NewTransactionRepository(), internal.NewStatementPrinter(logger))

	account.Deposit(*internal.NewTransaction(1000, "10-01-2012"))
	account.Deposit(*internal.NewTransaction(2000, "13-01-2012"))
	account.Withdrawal(*internal.NewTransaction(500, "14-01-2012"))

	account.PrintBalance()
}
