package banking_kata_test

import (
	banking_kata "github.com/adolsalamanca/banking-kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Banking kata tests", func () {

	type testableAccount struct {
		account * banking_kata.BankAccount
	}

	var bankAccount * banking_kata.BankAccount
	testAccount:= testableAccount{}

	BeforeEach(func() {
		bankAccount = banking_kata.NewBankAccount(banking_kata.NewTransactionRepository(), banking_kata.NewStatementPrinter(&log.Logger{}))
		testAccount.account = bankAccount
	})

	It("Should perform a store operation if requested", func() {

		transaction := banking_kata.NewTransaction(100,"11-09-2019")
		bankAccount.Store(*transaction)
		Expect(bankAccount.GetBalance()).To(BeEquivalentTo(100))
	})

	It("Should perform a withdrawal operation if requested", func() {

		transaction := banking_kata.NewTransaction(100,"11-09-2019")
		bankAccount.Withdrawal(*transaction)
		Expect(bankAccount.GetBalance()).To(BeEquivalentTo(-100))
	})

})
