package banking_kata_test

import (
	"github.com/adolsalamanca/banking-kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Banking kata tests", func () {

	It("Should retrieve a transaction properly created", func() {

		transaction := banking_kata.NewTransaction(200, "02-09-2019")
		Expect(transaction.Amount).To(BeEquivalentTo(200))
		Expect(transaction.Date).To(BeEquivalentTo("02-09-2019"))
	})

	It("Should retrieve a transaction repository properly created", func() {

		repository := banking_kata.NewTransactionRepository([]banking_kata.Transaction{
			*banking_kata.NewTransaction(200, "02-09-2019"),
			*banking_kata.NewTransaction(-100, "03-09-2019"),
		})

		Expect(len(repository.Transactions)).To(BeEquivalentTo(2))
	})
})
