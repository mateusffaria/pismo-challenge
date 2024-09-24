package repositories_test

import (
	"github.com/mateusffaria/pismo-challenge/internal/accounts/domains"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func NewAccountRepositoryMock() *AccountRepositoryMock {
	return &AccountRepositoryMock{}
}

func (m *AccountRepositoryMock) CreateUserAccount(acc domains.Account) (domains.Account, error) {
	args := m.Called(acc)

	return args.Get(0).(domains.Account), args.Error(1)
}

func (m *AccountRepositoryMock) GetUserAccount(id string) (domains.Account, error) {
	args := m.Called(id)

	return args.Get(0).(domains.Account), args.Error(1)
}
