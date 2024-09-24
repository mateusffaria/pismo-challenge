package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/domains"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/handlers/request"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/repositories"
	customErrors "github.com/mateusffaria/pismo-challenge/internal/transactions/services/errors"
)

type TransactionServiceProvider interface {
	CreateTransaction(tr request.NewTransactionRequest) (domains.Transaction, error)
}

type TransactionService struct {
	tr repositories.TransactionRepositoryProvider
}

func NewTransactionService(tr repositories.TransactionRepositoryProvider) *TransactionService {
	return &TransactionService{
		tr: tr,
	}
}

func (as *TransactionService) CreateTransaction(tr request.NewTransactionRequest) (domains.Transaction, error) {
	ac, err := as.tr.CreateTransaction(domains.Transaction{
		AccountId:       uuid.MustParse(tr.AccountId),
		OperationTypeId: tr.OperationTypeId,
		Amount:          tr.Amount,
	})
	if err != nil {
		log.Default().Printf("\nfailed to create user account %v\n", err)

		return ac, customErrors.ErrInternalDatabaseError
	}

	return ac, err
}
