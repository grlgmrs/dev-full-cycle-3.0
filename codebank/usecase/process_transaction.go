package usecase

import (
	"time"

	"github.com/codeedu/codebank/domain"
	"github.com/codeedu/codebank/dto"
)

type UseCaseTransaction struct {
	TransactionRepository domain.ITransactionRepository
}

func NewUseCaseTransaction(transactionRepository domain.ITransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

func (useCaseTransaction UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {
	creditCard := useCaseTransaction.HydrateCreditCard(transactionDto)
	ccBalanceAndLimit, err := useCaseTransaction.TransactionRepository.GetCreditCard(*creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}

	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance

	transaction := useCaseTransaction.NewTransaction(transactionDto, *creditCard)
	transaction.ProcessAndValidate(creditCard)
	
	err = useCaseTransaction.TransactionRepository.SaveTransaction(*transaction, *creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}

	return *transaction, nil
}

func (useCaseTransaction UseCaseTransaction) HydrateCreditCard(transactionDto dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDto.Name
	creditCard.Number = transactionDto.Number
	creditCard.ExpirationMonth = transactionDto.ExpirationMonth
	creditCard.ExpirationYear = transactionDto.ExpirationYear
	creditCard.CVV = transactionDto.CVV

	return creditCard
}

func (useCaseTransaction UseCaseTransaction) NewTransaction(transactionDto dto.Transaction, creditCard domain.CreditCard) *domain.Transaction {
	transaction := domain.NewTransaction()
	transaction.CreditCardId = creditCard.ID
	transaction.Amount = transactionDto.Amount
	transaction.Store = transactionDto.Store
	transaction.Description = transactionDto.Description
	transaction.CreatedAt = time.Now()

	return transaction
}