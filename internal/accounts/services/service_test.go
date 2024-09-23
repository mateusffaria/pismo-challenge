package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/domains"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request"
	repositories_test "github.com/mateusffaria/pismo-challenge/test/accounts/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldCreateUserAccountSuccessfully(t *testing.T) {
	repoMock := repositories_test.NewAccountRepositoryMock()
	accModel := domains.Account{
		ID:             uuid.New(),
		DocumentNumber: "1234567890",
	}
	repoMock.On("CreateUserAccount", mock.AnythingOfType("domains.Account")).Return(accModel, nil)

	svc := NewAccountService(repoMock)

	res, err := svc.CreateUserAccount(request.UserAccountRequest{
		DocumentNumber: "1234567890",
	})

	assert.NoError(t, err)
	assert.Equal(t, res, accModel)
}
