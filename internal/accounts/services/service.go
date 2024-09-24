package services

import (
	"errors"
	"log"

	"github.com/mateusffaria/pismo-challenge/internal/accounts/domains"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/repositories"
	customErrors "github.com/mateusffaria/pismo-challenge/internal/accounts/services/errors"
	"gorm.io/gorm"
)

type AccountServiceProvider interface {
	CreateUserAccount(uar request.UserAccountRequest) (domains.Account, error)
	GetUserAccount(id string) (domains.Account, error)
}

type AccountService struct {
	ar repositories.AccountRepositoryProvider
}

func NewAccountService(ar repositories.AccountRepositoryProvider) *AccountService {
	return &AccountService{
		ar: ar,
	}
}

func (as *AccountService) CreateUserAccount(uar request.UserAccountRequest) (domains.Account, error) {
	ac, err := as.ar.CreateUserAccount(domains.Account{
		DocumentNumber: uar.DocumentNumber,
	})
	if err != nil {
		log.Default().Printf("\nfailed to create user account %v\n", err)
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey):
			return ac, customErrors.ErrAccountDuplicated
		default:
			return ac, customErrors.ErrInternalDatabaseError
		}
	}

	return ac, err
}

func (as *AccountService) GetUserAccount(id string) (domains.Account, error) {
	ac, err := as.ar.GetUserAccount(id)
	if err != nil {
		log.Default().Printf("\nfailed to get user data %v\n", err)
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return ac, customErrors.ErrAccountNotFound
		default:
			return ac, customErrors.ErrInternalDatabaseError
		}
	}

	return ac, err
}
