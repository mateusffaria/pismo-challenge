package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/domains"
	"gorm.io/gorm"
)

type TransactionRepositoryProvider interface {
	CreateTransaction(acc domains.Transaction) (domains.Transaction, error)
}

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (ar *AccountRepository) CreateTransaction(tr domains.Transaction) (domains.Transaction, error) {
	tr.ID = uuid.New()
	tr.EventDate = time.Now().UTC()

	// TODO: Adjust to handle possible coner-cases beyond duplicate key
	if err := ar.DB.Create(&tr).Error; err != nil {
		return tr, err
	}

	return tr, nil
}
