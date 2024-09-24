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

	if err := ar.DB.Create(&acc).Error; err != nil {
		return acc, err
	}

	return acc, nil
}

func (ar *AccountRepository) GetUserAccount(id string) (domains.Account, error) {
	var acc domains.Account

	res := ar.DB.First(&acc, "id = ?", id)
	if res.Error != nil {
		return acc, res.Error
	}

	return acc, nil
}
