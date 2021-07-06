package repository

import (
	"database/sql"
	"errors"

	"github.com/codeedu/codebank/domain"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (transactionRepository *TransactionRepositoryDb) SaveTransaction(transaction domain.Transaction, creditCard domain.CreditCard) error {
	stmt, err :=
		transactionRepository.db.Prepare(
			`INSERT INTO transactions(id, credit_card_id, amount, status, description, store, created_at)
			VALUES($1, $2, $3, $4, $5, $6, $7)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreditCardId,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreatedAt,
	)

	if err != nil {
		return err
	}

	if transaction.Status == "approved" {
		err = transactionRepository.UpdateBalance(creditCard)

		if err != nil {
			return err
		}
	}
	err = stmt.Close()

	if err != nil {
		return err
	}

	return nil
}

func (transactionRepository *TransactionRepositoryDb) UpdateBalance(creditCard domain.CreditCard) error {
	_, err := transactionRepository.db.Exec(
		`UPDATE credit_cards SET balance = $1 WHERE id = $2`,
		creditCard.Balance, creditCard.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (transactionRepository *TransactionRepositoryDb) CreateCreditCard(creditCard domain.CreditCard) error {
	stmt, err := transactionRepository.db.Prepare(
		`INSERT INTO credit_cards(id, name, number, expiration_month, expiration_year, cvv, balance, balance_limit)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
	)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.Name,
		creditCard.Number,
		creditCard.ExpirationMonth,
		creditCard.ExpirationYear,
		creditCard.CVV,
		creditCard.Balance,
		creditCard.Limit,
	)

	if err != nil {
		return err
	}

	err = stmt.Close()

	if err != nil {
		return err
	}

	return nil
}

func (transactionRepository *TransactionRepositoryDb) GetCreditCard(creditCard domain.CreditCard) (domain.CreditCard, error) {
	var newCreditCard domain.CreditCard
	stmt, err := transactionRepository.db.Prepare("SELECT id, balance, balance_limit, FROM credit_cards WHERE number=$1")
	if err != nil {
		return newCreditCard, err
	}

	if err = stmt.QueryRow(creditCard.Number).Scan(
		&newCreditCard.ID, 
		&newCreditCard.Balance,
		&newCreditCard.Limit,
	); err != nil {
		return newCreditCard, errors.New("credit card does not exists")
	}

	return newCreditCard, nil
}