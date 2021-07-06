package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codeedu/codebank/domain"
	"github.com/codeedu/codebank/infrastructure/repository"
	"github.com/codeedu/codebank/usecase"
	_ "github.com/lib/pq" /// Escrever dessa forma permitirá que use métodos do pacote, sem precisar especificá-lo
)

func main() {
	db :=  setupDb()
	/// Roda por último na função, então pode-se dizer que seria a mesma coisa que deixar na última linha
	defer db.Close()

	creditCard := domain.NewCreditCard()
	creditCard.Number = "4321"
	creditCard.Name = "Gabriel"
	creditCard.ExpirationMonth = 7
	creditCard.ExpirationYear = 22
	creditCard.Limit = 2500
	creditCard.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*creditCard)
	if err != nil {
		fmt.Println(err)
	}
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transationRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transationRepository)

	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}

	return db
}