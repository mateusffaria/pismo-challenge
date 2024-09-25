package repositories

import (
	"time"

	"pismo-challenge/internal/transactions/domains"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepositoryProvider interface {
	CreateTransaction(acc domains.Transaction) (domains.Transaction, error)
}

type AccountRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (ar *AccountRepository) CreateTransaction(tr domains.Transaction) (domains.Transaction, error) {
	tr.ID = uuid.New()
	tr.EventDate = time.Now().UTC()

	if err := ar.DB.Create(&tr).Error; err != nil {
		return tr, err
	}

	return tr, nil
}
