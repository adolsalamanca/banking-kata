package banking_kata

type Transaction struct {
	Amount float32
	Date string
}

func NewTransaction(amount float32, date string) *Transaction{
	return &Transaction{Amount: amount, Date: date}
}

type TransactionRepository struct {
	Transactions [] Transaction
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) AddTransaction(transaction Transaction) {
	r.Transactions = append(r.Transactions, transaction)
}
