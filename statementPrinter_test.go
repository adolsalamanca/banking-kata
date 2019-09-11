package banking_kata_test

import (
	"bytes"
	"github.com/adolsalamanca/banking-kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Banking kata tests", func () {

	It("Should print statement headers if requested", func() {

		var buf bytes.Buffer
		loggerToTest := log.Logger{}
		loggerToTest.SetOutput(&buf)

		statementPrinter := banking_kata.NewStatementPrinter(&loggerToTest)
		statementPrinter.PrintStatementHeaders()

		headers := "date || credit || debit || balance\n"
		Expect(buf.String()).To(BeEquivalentTo(headers))
	})

	It("Should print statements and headers if requested", func() {

		var buf bytes.Buffer
		loggerToTest := log.Logger{}
		loggerToTest.SetOutput(&buf)

		statementPrinter := banking_kata.NewStatementPrinter(&loggerToTest)
		transaction1 := banking_kata.NewTransaction(500,"10-09-2019")
		transaction2 := banking_kata.NewTransaction(-100,"11-09-2019")

		line1 := banking_kata.NewStatementLine(*transaction1, 500)
		line2 := banking_kata.NewStatementLine(*transaction2, 400)

		lines := []banking_kata.StatementLine{line1,line2}

		statementPrinter.PrintStatement(lines)

		headers := "date || credit || debit || balance\n"
		textString1 := "11-09-2019 || || 100 || 400\n"
		textString2 := "10-09-2019 || 500 || || 500\n"

		Expect(buf.String()).To(BeEquivalentTo(headers+textString1+textString2))
	})

})
