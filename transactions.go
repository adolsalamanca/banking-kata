package banking_kata

type Transaction struct {
	Amount float32
	Date string
}


func NewTransaction(amount float32, date string) *Transaction{
	return &Transaction{Amount: amount, Date: date}
}

