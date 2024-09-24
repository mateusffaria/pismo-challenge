package services_test

import (
	"github.com/mateusffaria/pismo-challenge/internal/accounts/domains"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request"
	"github.com/stretchr/testify/mock"
)

type AccountsServiceRepository struct {
	mock.Mock
}

func NewAccountsService() *AccountsServiceRepository {
	return &AccountsServiceRepository{}
}

// CreateUserAccount implements services.AccountServiceProvider.
func (m *AccountsServiceRepository) CreateUserAccount(uar request.UserAccountRequest) (domains.Account, error) {
	args := m.Called(uar)

	return args.Get(0).(domains.Account), args.Error(1)
}

func (m *AccountsServiceRepository) GetUserAccount(id string) (domains.Account, error) {
	args := m.Called(id)

	return args.Get(0).(domains.Account), args.Error(1)
}
