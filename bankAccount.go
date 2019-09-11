package banking_kata


type BankAccount struct {
	repository       *TransactionRepository
	statementPrinter *StatementPrinter
}

func (account *BankAccount) Store(t Transaction) {
	account.repository.Transactions = append(account.repository.Transactions, t)
}

func (account *BankAccount) Withdrawal(t Transaction) {
	t.Amount = -t.Amount
	account.repository.Transactions = append(account.repository.Transactions, t)
}

func (account *BankAccount) GetBalance() float32 {
	total := float32(0)
	for _, val := range account.repository.Transactions {
		total += val.Amount
	}
	return total
}

func NewBankAccount(r *TransactionRepository, sp *StatementPrinter) *BankAccount {
	return &BankAccount{repository: r, statementPrinter: sp}
}
