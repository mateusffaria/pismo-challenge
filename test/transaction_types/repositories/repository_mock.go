package repositories_test

import (
	"github.com/mateusffaria/pismo-challenge/internal/transaction_types/domains"
	"github.com/stretchr/testify/mock"
)

type TransactionTypesRepository struct {
	mock.Mock
}

func NewTransactionTypesRepository() *TransactionTypesRepository {
	return &TransactionTypesRepository{}
}

func (m *TransactionTypesRepository) GetTransactionTypeById(id int) (domains.TransactionType, error) {
	args := m.Called(id)

	return args.Get(0).(domains.TransactionType), args.Error(1)
}
