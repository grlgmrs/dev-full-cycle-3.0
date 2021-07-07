package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/codeedu/codebank/domain"
	"github.com/codeedu/codebank/infrastructure/grpc/server"
	"github.com/codeedu/codebank/infrastructure/kafka"
	"github.com/codeedu/codebank/infrastructure/repository"
	"github.com/codeedu/codebank/usecase"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" /// Escrever dessa forma permitirá que use métodos do pacote, sem precisar especificá-lo
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	db := SetupDb()
	/// Roda por último na função, então pode-se dizer que seria a mesma coisa que deixar na última linha
	defer db.Close()

	producer := SetupKafkaProducer()
	processTransactionUseCase := SetupTransactionUseCase(db, producer)
	ServeGrpc(processTransactionUseCase)
}

func SetupTransactionUseCase(db *sql.DB, producer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transationRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transationRepository)
	useCase.KafkaProducer = producer

	return useCase
}

func SetupDb() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"),
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}

	return db
}

func SetupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer(os.Getenv("KafkaBootstrapServers"))

	return producer
}

func ServeGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	fmt.Println("Rodando gRPC Server")
	grpcServer.Serve()
}

func CreateFakeCreditCard(db *sql.DB) {
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
