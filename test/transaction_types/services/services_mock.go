package services_test

import (
	"github.com/mateusffaria/pismo-challenge/internal/transaction_types/domains"
	"github.com/stretchr/testify/mock"
)

type TransactionTypesServiceRepository struct {
	mock.Mock
}

func NewTransactionTypesService() *TransactionTypesServiceRepository {
	return &TransactionTypesServiceRepository{}
}

func (m *TransactionTypesServiceRepository) GetTransactionType(transactionTypeId int) (domains.TransactionType, error) {
	args := m.Called(transactionTypeId)

	return args.Get(0).(domains.TransactionType), args.Error(1)
}
