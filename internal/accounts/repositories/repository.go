package repositories

import (
	"github.com/google/uuid"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/domains"
	"gorm.io/gorm"
)

type AccountRepositoryProvider interface {
	CreateUserAccount(acc domains.Account) (domains.Account, error)
	GetUserAccount(id string) (domains.Account, error)
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
	acc.ID = uuid.New()

	// TODO: Adjust to handle possible coner-cases beyond duplicate key
	if err := ar.DB.Create(&acc).Error; err != nil {
		return acc, err
	}

	return acc, nil
}

func (ar *AccountRepository) GetUserAccount(id string) (domains.Account, error) {
	var ua domains.Account
	res := ar.DB.First(&ua, "id = ?", id)
	if res.Error != nil {
		return ua, res.Error
	}

	return ua, nil
}
