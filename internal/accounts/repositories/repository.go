package repositories

import (
	"github.com/google/uuid"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/domains"
	"gorm.io/gorm"
)

type AccountRepositoryProvider interface {
	CreateUserAccount(acc domains.Account) (domains.Account, error)
}

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (ar *AccountRepository) CreateUserAccount(acc domains.Account) (domains.Account, error) {
	acc.ID = uuid.New().String()

	err := ar.DB.Create(&acc).Error
	if err != nil {
		return acc, err
	}

	return acc, nil
}
