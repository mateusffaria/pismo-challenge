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
	ar repositories.AccountRepositoryProvider
}

func NewAccountService(ar repositories.AccountRepositoryProvider) *AccountService {
	return &AccountService{
		ar: ar,
	}
}

func (as *AccountService) CreateUserAccount(uar request.UserAccountRequest) (domains.Account, error) {

	return as.ar.CreateUserAccount(domains.Account{
		DocumentNumber: uar.DocumentNumber,
	})
}
