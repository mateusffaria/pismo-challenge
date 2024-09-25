package services_test

import (
	"pismo-challenge/internal/accounts/domains"
	"pismo-challenge/internal/accounts/handlers/request"

	"github.com/stretchr/testify/mock"
)

type AccountsServiceRepository struct {
	mock.Mock
}

func NewAccountsService() *AccountsServiceRepository {
	return &AccountsServiceRepository{}
}

// CreateAccount implements services.AccountServiceProvider.
func (m *AccountsServiceRepository) CreateAccount(uar request.AccountRequest) (domains.Account, error) {
	args := m.Called(uar)

	return args.Get(0).(domains.Account), args.Error(1)
}

func (m *AccountsServiceRepository) GetAccount(id string) (domains.Account, error) {
	args := m.Called(id)

	return args.Get(0).(domains.Account), args.Error(1)
}
