package banking_kata_test

import (
	"github.com/adolsalamanca/banking-kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Banking kata tests", func () {

	It("Should generate bank statement after a few operations", func() {

		transaction := banking_kata.NewTransaction(200, "02-09-2019")
		Expect(transaction.Amount).To(BeEquivalentTo(200))
		Expect(transaction.Date).To(BeEquivalentTo("02-09-2019"))

	})

})
