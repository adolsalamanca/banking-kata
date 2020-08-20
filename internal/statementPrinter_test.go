package internal_test

import (
	"bytes"
	"github.com/adolsalamanca/banking-kata/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Banking kata tests", func() {

	It("Should print statements and headers if requested", func() {

		var buf bytes.Buffer
		loggerToTest := log.Logger{}
		loggerToTest.SetOutput(&buf)

		statementPrinter := internal.NewStatementPrinter(&loggerToTest)
		transaction1 := internal.NewTransaction(500, "10-09-2019")
		transaction2 := internal.NewTransaction(-100, "11-09-2019")

		line1 := internal.NewStatementLine(*transaction1, 500)
		line2 := internal.NewStatementLine(*transaction2, 400)

		lines := []internal.StatementLine{line1, line2}

		statementPrinter.PrintAllStatements(lines)

		headers := "date || credit || debit || balance\n"
		textString1 := "11-09-2019 || || 100 || 400\n"
		textString2 := "10-09-2019 || 500 || || 500\n"

		Expect(buf.String()).To(BeEquivalentTo(headers + textString1 + textString2))
	})

})
