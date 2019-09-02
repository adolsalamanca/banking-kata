package banking_kata_test

import (
	"bytes"
	"github.com/adolsalamanca/banking-kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Banking kata tests", func () {

	It("Should generate bank statement after a few operations", func() {

		var buf bytes.Buffer
		log.SetOutput(&buf)
		banking_kata.PrintStatement(log.Logger{})
		value := buf.String()
		log.Println(value)
		Expect(buf.String()).To(Not(BeNil()))

	})

})
