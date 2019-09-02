package banking_kata

type TransactionRepository struct {
	Transactions [] Transaction
}

func NewTransactionRepository(transactions [] Transaction) *TransactionRepository {
	return &TransactionRepository{Transactions:transactions}
}

type Transaction struct {
	Amount float32
	Date string
}

func NewTransaction(amount float32, date string) *Transaction{
	return &Transaction{Amount: amount, Date: date}
}

