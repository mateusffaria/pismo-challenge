package services

import (
	"log"

	accSvc "pismo-challenge/internal/accounts/services"
	ttSvc "pismo-challenge/internal/operation_types/services"
	"pismo-challenge/internal/transactions/domains"
	"pismo-challenge/internal/transactions/handlers/request"
	"pismo-challenge/internal/transactions/repositories"
	customErrors "pismo-challenge/internal/transactions/services/errors"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionServiceProvider interface {
	CreateTransaction(tr request.NewTransactionRequest) (domains.Transaction, error)
}

type TransactionService struct {
	tr repositories.TransactionRepositoryProvider
	as accSvc.AccountServiceProvider
	tt ttSvc.OperationTypesServiceProvider
}

func NewTransactionService(
	tr repositories.TransactionRepositoryProvider,
	as accSvc.AccountServiceProvider,
	tt ttSvc.OperationTypesServiceProvider,
) *TransactionService {
	return &TransactionService{
		tr: tr,
		as: as,
		tt: tt,
	}
}

func (ts *TransactionService) CreateTransaction(tr request.NewTransactionRequest) (domains.Transaction, error) {
	_, err := ts.as.GetAccount(tr.AccountId)
	if err != nil {
		log.Default().Printf("\nfailed getting account owner info %v\n", err)
		return domains.Transaction{}, err
	}

	_, err = ts.tt.GetOperationType(tr.OperationTypeId)
	if err != nil {
		log.Default().Printf("\nfailed getting operation type %v\n", err)

		return domains.Transaction{}, err
	}

	ac, err := ts.tr.CreateTransaction(domains.Transaction{
		AccountId:       uuid.MustParse(tr.AccountId),
		OperationTypeId: tr.OperationTypeId,
		Amount:          decimal.NewFromFloat(tr.Amount),
	})
	if err != nil {
		log.Default().Printf("\nfailed to create transaction %v\n", err)

		return ac, customErrors.ErrInternalDatabaseError
	}

	return ac, err
}
