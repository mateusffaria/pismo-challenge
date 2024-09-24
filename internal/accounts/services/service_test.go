package services

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/domains"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request"
	svcErrors "github.com/mateusffaria/pismo-challenge/internal/accounts/services/errors"
	repositories_test "github.com/mateusffaria/pismo-challenge/test/accounts/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestShouldCreateAccountSuccessfully(t *testing.T) {
	repoMock := repositories_test.NewAccountRepositoryMock()
	accModel := domains.Account{
		ID:             uuid.New(),
		DocumentNumber: "1234567890",
	}
	repoMock.On("CreateAccount", mock.AnythingOfType("domains.Account")).Return(accModel, nil)

	svc := NewAccountService(repoMock)

	res, err := svc.CreateAccount(request.AccountRequest{
		DocumentNumber: "1234567890",
	})

	assert.NoError(t, err)
	assert.Equal(t, res, accModel)
}

func TestShouldHandleAccountDuplicateErrors(t *testing.T) {
	repoMock := repositories_test.NewAccountRepositoryMock()
	repoMock.On("CreateAccount", mock.AnythingOfType("domains.Account")).Return(domains.Account{}, gorm.ErrDuplicatedKey)

	svc := NewAccountService(repoMock)

	_, err := svc.CreateAccount(request.AccountRequest{
		DocumentNumber: "1234567890",
	})

	assert.ErrorIs(t, err, svcErrors.ErrAccountDuplicated)
}

func TestShouldHandleAccountCreationErrors(t *testing.T) {
	repoMock := repositories_test.NewAccountRepositoryMock()
	repoMock.On("CreateAccount", mock.AnythingOfType("domains.Account")).Return(domains.Account{}, fmt.Errorf("some error"))

	svc := NewAccountService(repoMock)

	_, err := svc.CreateAccount(request.AccountRequest{
		DocumentNumber: "1234567890",
	})

	assert.ErrorIs(t, err, svcErrors.ErrInternalDatabaseError)
}

func TestShouldGetAccountSuccessfully(t *testing.T) {
	repoMock := repositories_test.NewAccountRepositoryMock()
	accModel := domains.Account{
		ID:             uuid.New(),
		DocumentNumber: "1234567890",
	}
	repoMock.On("GetAccount", "1234567890").Return(accModel, nil)

	svc := NewAccountService(repoMock)

	res, err := svc.GetAccount("1234567890")

	assert.NoError(t, err)
	assert.Equal(t, res, accModel)
}

func TestShouldHandleAccountNotFound(t *testing.T) {
	repoMock := repositories_test.NewAccountRepositoryMock()
	repoMock.On("GetAccount", "1234567890").Return(domains.Account{}, gorm.ErrRecordNotFound)

	svc := NewAccountService(repoMock)

	_, err := svc.GetAccount("1234567890")

	assert.ErrorIs(t, err, svcErrors.ErrAccountNotFound)
}

func TestShouldHandleGetAccountErrors(t *testing.T) {
	repoMock := repositories_test.NewAccountRepositoryMock()
	repoMock.On("GetAccount", "1234567890").Return(domains.Account{}, fmt.Errorf("some error"))

	svc := NewAccountService(repoMock)

	_, err := svc.GetAccount("1234567890")

	assert.ErrorIs(t, err, svcErrors.ErrInternalDatabaseError)
}
