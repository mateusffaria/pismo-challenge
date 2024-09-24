package repositories

import (
	"github.com/mateusffaria/pismo-challenge/internal/transaction_types/domains"
	"gorm.io/gorm"
)

type TransactionTypeRepositoryProvider interface {
	GetTransactionTypeById(id int) (domains.TransactionType, error)
}

type TransactionType struct {
	DB *gorm.DB
}

func NewTransactionType(db *gorm.DB) *TransactionType {
	return &TransactionType{
		DB: db,
	}
}

func (ar *TransactionType) GetTransactionTypeById(id int) (domains.TransactionType, error) {
	var tt domains.TransactionType

	res := ar.DB.First(&tt, "id = ?", id)
	if res.Error != nil {
		return tt, res.Error
	}

	return tt, nil
}
