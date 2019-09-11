package banking_kata_test

import (
	"bytes"
	banking_kata "github.com/adolsalamanca/banking-kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Banking kata tests", func () {

	var account *banking_kata.Account
	var logger *log.Logger
	var buffer bytes.Buffer

	BeforeEach(func() {
		logger = &log.Logger{}
		buffer = bytes.Buffer{}

		logger.SetOutput(&buffer)
		account = banking_kata.NewBankAccount(banking_kata.NewTransactionRepository(), banking_kata.NewStatementPrinter(logger))
	})

	It("Should perform a store operation if requested", func() {

		transaction := banking_kata.NewTransaction(100,"11-09-2019")
		account.Store(*transaction)
		account.PrintBalance()

		Expect(buffer.String()).To(ContainSubstring("11-09-2019 || 100"))
	})

	It("Should perform a withdrawal operation if requested", func() {

		transaction := banking_kata.NewTransaction(100,"11-09-2019")
		account.Withdrawal(*transaction)
		account.PrintBalance()

		Expect(buffer.String()).To(ContainSubstring("11-09-2019 || || 100"))
	})


	It("Should print statement properly after a few transactions", func() {

		transaction1 := banking_kata.NewTransaction(100,"10-09-2019")
		transaction2 := banking_kata.NewTransaction(50,"11-09-2019")

		account.Store(*transaction1)
		account.Withdrawal(*transaction2)

		headers := "date || credit || debit || balance\n"
		textString1 := "11-09-2019 || || 50 || 50\n"
		textString2 := "10-09-2019 || 100 || || 100\n"

		account.PrintBalance()
		Expect(buffer.String()).To(BeEquivalentTo(headers+textString1+textString2))
	})

})
