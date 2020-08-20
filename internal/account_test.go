package internal_test

import (
	"bytes"
	"github.com/adolsalamanca/banking-kata/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Banking kata tests", func() {

	var account *internal.Account
	var logger *log.Logger
	var buffer bytes.Buffer

	BeforeEach(func() {
		logger = &log.Logger{}
		buffer = bytes.Buffer{}

		logger.SetOutput(&buffer)
		account = internal.NewAccount(internal.NewTransactionRepository(), internal.NewStatementPrinter(logger))
	})

	It("Should perform a deposit operation if requested", func() {

		transaction := internal.NewTransaction(100, "11-09-2019")
		account.Deposit(*transaction)
		account.PrintBalance()

		Expect(buffer.String()).To(ContainSubstring("11-09-2019 || 100"))
	})

	It("Should perform a withdrawal operation if requested", func() {

		transaction := internal.NewTransaction(100, "11-09-2019")
		account.Withdrawal(*transaction)
		account.PrintBalance()

		Expect(buffer.String()).To(ContainSubstring("11-09-2019 || || 100"))
	})

	It("Should print statement properly after a few transactions", func() {

		transaction1 := internal.NewTransaction(100, "10-09-2019")
		transaction2 := internal.NewTransaction(50, "11-09-2019")

		account.Deposit(*transaction1)
		account.Withdrawal(*transaction2)

		headers := "date || credit || debit || balance\n"
		textString1 := "11-09-2019 || || 50 || 50\n"
		textString2 := "10-09-2019 || 100 || || 100\n"

		account.PrintBalance()
		Expect(buffer.String()).To(BeEquivalentTo(headers + textString1 + textString2))
	})

})
