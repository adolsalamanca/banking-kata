package banking_kata


type BankAccount struct {
	repository       *TransactionRepository
	statementPrinter *StatementPrinter
}

func NewBankAccount(r *TransactionRepository, sp *StatementPrinter) *BankAccount {
	return &BankAccount{repository: r, statementPrinter: sp}
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

func (account *BankAccount) PrintBalance() {
	statementLines := account.retrieveStatementLines()
	account.statementPrinter.PrintStatement(statementLines)
}

func (account *BankAccount) retrieveStatementLines() []StatementLine{
	var total float32 = 0.0
	var statementLines []StatementLine
	for _, transaction := range account.repository.Transactions {
		total += transaction.Amount
		statementLines = append(statementLines, NewStatementLine(transaction, total))
	}
	return statementLines
}