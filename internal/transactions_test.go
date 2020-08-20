package internal_test

import (
	"github.com/adolsalamanca/banking-kata/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Banking kata tests", func() {

	It("Should retrieve a transaction properly created", func() {

		transaction := internal.NewTransaction(200, "02-09-2019")
		Expect(transaction.Amount).To(BeEquivalentTo(200))
		Expect(transaction.Date).To(BeEquivalentTo("02-09-2019"))
	})

	It("Should retrieve an empty transaction repository properly created", func() {

		repository := internal.NewTransactionRepository()
		Expect(len(repository.Transactions)).To(BeEquivalentTo(0))
	})

	It("Should store transactions if add transaction is called", func() {

		repository := internal.NewTransactionRepository()
		repository.AddTransaction(*internal.NewTransaction(200, "03-09-2019"))
		Expect(len(repository.Transactions)).To(BeEquivalentTo(1))
	})
})
