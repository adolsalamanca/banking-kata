package banking_kata


type BankAccount struct {
	repository       *TransactionRepository
	statementPrinter *StatementPrinter
	Balance float32
}

func (account *BankAccount) Store(t Transaction) {
	account.repository.Transactions = append(account.repository.Transactions, t)
	account.Balance += t.Amount
}

func NewBankAccount(r *TransactionRepository, sp *StatementPrinter) *BankAccount {
	return &BankAccount{repository: r, statementPrinter: sp}
}
