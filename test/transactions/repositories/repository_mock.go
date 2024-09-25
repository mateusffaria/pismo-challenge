package repositories_test

import (
	"pismo-challenge/internal/transactions/domains"

	"github.com/stretchr/testify/mock"
)

type TransactionRepository struct {
	mock.Mock
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (m *TransactionRepository) CreateTransaction(acc domains.Transaction) (domains.Transaction, error) {
	args := m.Called(acc)

	return args.Get(0).(domains.Transaction), args.Error(1)
}
