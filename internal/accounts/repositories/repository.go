package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/domains"
	customError "github.com/mateusffaria/pismo-challenge/internal/accounts/repositories/errors"
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
	acc.ID = uuid.New()

	if err := ar.DB.Create(&acc).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
		return acc, customError.NewDuplicateEntity()
	}

	return acc, nil
}
