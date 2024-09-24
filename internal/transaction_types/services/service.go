package services

import (
	"github.com/mateusffaria/pismo-challenge/internal/transaction_types/domains"
	"github.com/mateusffaria/pismo-challenge/internal/transaction_types/repositories"
)

type TransactionTypesServiceProvider interface {
	GetTransactionType(transactionTypeId int) (domains.TransactionType, error)
}

type TransactionTypesService struct {
	tr repositories.TransactionTypeRepositoryProvider
}

func NewTransactionTypesService(tr repositories.TransactionTypeRepositoryProvider) *TransactionTypesService {
	return &TransactionTypesService{
		tr: tr,
	}
}

func (ts *TransactionTypesService) GetTransactionType(transactionTypeId int) (domains.TransactionType, error) {
	tt, err := ts.tr.GetTransactionTypeById(transactionTypeId)
	if err != nil {
		return tt, err
	}

	return tt, nil
}
