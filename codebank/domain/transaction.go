package domain

import (
	"time"
	uuid "github.com/satori/go.uuid"
)

type ITransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCard(creditCard CreditCard) (CreditCard, error)
	CreateCreditCard(creditCard CreditCard) error
}

type Transaction struct {
	ID string
	Amount float64
	Status string
	Description string
	Store string
	CreditCardId string
	CreatedAt time.Time
}

/// Igual ao C, o * aqui representa um ponteiro
func NewTransaction() *Transaction {
	/// Como estou me referindo à um ponteiro, naturalmente para o acessarmos, necessitaremos 
	/// pegar o seu endereço de memória, portanto, utilizando o &
	transaction := &Transaction{}
	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	return transaction
}

func (transaction *Transaction) ProcessAndValidate(creditCard *CreditCard) {
	if transaction.Amount + creditCard.Balance > creditCard.Limit {
		transaction.Status = "rejected"
	} else {
		transaction.Status = "approved"
		creditCard.Balance += transaction.Amount
	}
}