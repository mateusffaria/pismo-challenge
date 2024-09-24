package services

import (
	"log"

	"github.com/google/uuid"
	accSvc "github.com/mateusffaria/pismo-challenge/internal/accounts/services"
	ttSvc "github.com/mateusffaria/pismo-challenge/internal/transaction_types/services"
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
	as accSvc.AccountServiceProvider
	tt ttSvc.TransactionTypesServiceProvider
}

func NewTransactionService(
	tr repositories.TransactionRepositoryProvider,
	as accSvc.AccountServiceProvider,
	tt ttSvc.TransactionTypesServiceProvider,
) *TransactionService {
	return &TransactionService{
		tr: tr,
		as: as,
		tt: tt,
	}
}

func (ts *TransactionService) CreateTransaction(tr request.NewTransactionRequest) (domains.Transaction, error) {
	_, err := ts.as.GetUserAccount(tr.AccountId)
	if err != nil {
		log.Default().Printf("\nfailed getting account owner info %v\n", err)
		return domains.Transaction{}, err
	}

	_, err = ts.tt.GetTransactionType(tr.OperationTypeId)
	if err != nil {
		log.Default().Printf("\nfailed getting operation type %v\n", err)

		return domains.Transaction{}, err
	}

	ac, err := ts.tr.CreateTransaction(domains.Transaction{
		AccountId:       uuid.MustParse(tr.AccountId),
		OperationTypeId: tr.OperationTypeId,
		Amount:          tr.Amount,
	})
	if err != nil {
		log.Default().Printf("\nfailed to create transaction %v\n", err)

		return ac, customErrors.ErrInternalDatabaseError
	}

	return ac, err
}
