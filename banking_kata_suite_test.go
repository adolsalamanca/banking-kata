package banking_kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBankingKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BankingKata Suite")
}
