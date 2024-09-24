package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/domains"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/handlers/request"
	svcErrors "github.com/mateusffaria/pismo-challenge/internal/transactions/services/errors"
	repositories_test "github.com/mateusffaria/pismo-challenge/test/transactions/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldCreateUserAccountSuccessfully(t *testing.T) {
	repoMock := repositories_test.NewTransactionRepository()
	accModel := domains.Transaction{
		ID:              uuid.New(),
		AccountId:       uuid.New(),
		OperationTypeId: 1,
		Amount:          123.4,
		EventDate:       time.Now(),
	}
	repoMock.On("CreateTransaction", mock.AnythingOfType("domains.Transaction")).Return(accModel, nil)

	svc := NewTransactionService(repoMock)

	res, err := svc.CreateTransaction(request.NewTransactionRequest{
		AccountId:       "",
		OperationTypeId: 1,
		Amount:          123.4,
	})

	assert.NoError(t, err)
	assert.Equal(t, res, accModel)
}

func TestShouldHandleAccountCreationErrors(t *testing.T) {
	repoMock := repositories_test.NewTransactionRepository()
	repoMock.On("CreateTransaction", mock.AnythingOfType("domains.Transaction")).Return(domains.Transaction{}, fmt.Errorf("some error"))

	svc := NewTransactionService(repoMock)

	_, err := svc.CreateTransaction(request.NewTransactionRequest{
		AccountId:       "",
		OperationTypeId: 1,
		Amount:          123.4,
	})

	assert.ErrorIs(t, err, svcErrors.ErrInternalDatabaseError)
}
