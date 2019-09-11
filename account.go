package banking_kata


type Account struct {
	repository       *TransactionRepository
	statementPrinter *StatementPrinter
}

func NewBankAccount(r *TransactionRepository, sp *StatementPrinter) *Account {
	return &Account{repository: r, statementPrinter: sp}
}

func (a *Account) Store(t Transaction) {
	a.repository.Transactions = append(a.repository.Transactions, t)
}

func (a *Account) Withdrawal(t Transaction) {
	t.Amount = -t.Amount
	a.repository.Transactions = append(a.repository.Transactions, t)
}

func (a *Account) PrintBalance() {
	statementLines := a.retrieveStatementLines()
	a.statementPrinter.PrintStatement(statementLines)
}

func (a *Account) retrieveStatementLines() []StatementLine{
	var total float32 = 0.0
	var statementLines []StatementLine
	for _, transaction := range a.repository.Transactions {
		total += transaction.Amount
		statementLines = append(statementLines, NewStatementLine(transaction, total))
	}
	return statementLines
}