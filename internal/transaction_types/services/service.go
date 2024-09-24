package services

import (
	"errors"

	"github.com/mateusffaria/pismo-challenge/internal/transaction_types/domains"
	"github.com/mateusffaria/pismo-challenge/internal/transaction_types/repositories"
	customErrors "github.com/mateusffaria/pismo-challenge/internal/transaction_types/services/errors"
	"gorm.io/gorm"
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
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return tt, customErrors.ErrNotFound
		default:
			return tt, customErrors.ErrInternalDatabaseError
		}
	}

	return tt, nil
}
