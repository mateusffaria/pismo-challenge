package services

import (
	"github.com/mateusffaria/pismo-challenge/internal/accounts/domains"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/repositories"
)

type AccountServiceProvider interface {
	CreateUserAccount(uar request.UserAccountRequest) (domains.Account, error)
}

type AccountService struct {
	AccountRepository repositories.AccountRepositoryProvider
}

func NewAccountService(ar repositories.AccountRepositoryProvider) *AccountService {
	return &AccountService{
		AccountRepository: ar,
	}
}

func (as *AccountService) CreateUserAccount(uar request.UserAccountRequest) (domains.Account, error) {

	return as.AccountRepository.CreateUserAccount(domains.Account{
		DocumentNumber: uar.DocumentNumber,
	})
}
