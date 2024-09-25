package services

import (
	"fmt"
	"testing"
	"time"

	accDomain "pismo-challenge/internal/accounts/domains"
	ttDomain "pismo-challenge/internal/operation_types/domains"
	"pismo-challenge/internal/transactions/domains"
	"pismo-challenge/internal/transactions/handlers/request"
	svcErrors "pismo-challenge/internal/transactions/services/errors"
	accRepo "pismo-challenge/test/accounts/services"
	ttSvc "pismo-challenge/test/operation_types/services"
	repositories_test "pismo-challenge/test/transactions/repositories"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldCreateTransactionSuccessfully(t *testing.T) {
	accId := uuid.New()

	accSvc := accRepo.NewAccountsService()
	accSvc.On("GetAccount", accId.String()).Return(accDomain.Account{}, nil)

	ttSvc := ttSvc.NewOperationTypesService()
	ttSvc.On("GetOperationType", 1).Return(ttDomain.OperationType{}, nil)

	accModel := domains.Transaction{
		ID:              uuid.New(),
		AccountId:       uuid.New(),
		OperationTypeId: 1,
		Amount:          decimal.NewFromFloat(123.4),
		EventDate:       time.Now(),
	}

	repoMock := repositories_test.NewTransactionRepository()
	repoMock.On("CreateTransaction", mock.AnythingOfType("domains.Transaction")).Return(accModel, nil)

	svc := NewTransactionService(repoMock, accSvc, ttSvc)

	res, err := svc.CreateTransaction(request.NewTransactionRequest{
		AccountId:       accId.String(),
		OperationTypeId: 1,
		Amount:          123.4,
	})

	assert.NoError(t, err)
	assert.Equal(t, res, accModel)
}

func TestShouldHandleTransactionCreationErrors(t *testing.T) {
	accId := uuid.New()

	accSvc := accRepo.NewAccountsService()
	accSvc.On("GetAccount", accId.String()).Return(accDomain.Account{}, nil)

	ttSvc := ttSvc.NewOperationTypesService()
	ttSvc.On("GetOperationType", 1).Return(ttDomain.OperationType{}, nil)

	repoMock := repositories_test.NewTransactionRepository()
	repoMock.On("CreateTransaction", mock.AnythingOfType("domains.Transaction")).Return(domains.Transaction{}, fmt.Errorf("some error"))

	svc := NewTransactionService(repoMock, accSvc, ttSvc)

	_, err := svc.CreateTransaction(request.NewTransactionRequest{
		AccountId:       accId.String(),
		OperationTypeId: 1,
		Amount:          123.4,
	})

	assert.ErrorIs(t, err, svcErrors.ErrInternalDatabaseError)
}

func TestShouldHandleAccountInfoErrors(t *testing.T) {
	accId := uuid.New()

	accSvc := accRepo.NewAccountsService()
	accSvc.On("GetAccount", accId.String()).Return(accDomain.Account{}, fmt.Errorf("some error"))

	ttSvc := ttSvc.NewOperationTypesService()
	ttSvc.On("GetOperationType", 1).Return(ttDomain.OperationType{}, nil)

	repoMock := repositories_test.NewTransactionRepository()
	repoMock.On("CreateTransaction", mock.AnythingOfType("domains.Transaction")).Return(domains.Transaction{}, fmt.Errorf("some error"))

	svc := NewTransactionService(repoMock, accSvc, ttSvc)

	_, err := svc.CreateTransaction(request.NewTransactionRequest{
		AccountId:       accId.String(),
		OperationTypeId: 1,
		Amount:          123.4,
	})

	assert.ErrorContains(t, err, "some error")
}

func TestShouldHandleOperationTypeInfoErrors(t *testing.T) {
	accId := uuid.New()

	accSvc := accRepo.NewAccountsService()
	accSvc.On("GetAccount", accId.String()).Return(accDomain.Account{}, nil)

	ttSvc := ttSvc.NewOperationTypesService()
	ttSvc.On("GetOperationType", 1).Return(ttDomain.OperationType{}, fmt.Errorf("some error"))

	repoMock := repositories_test.NewTransactionRepository()
	repoMock.On("CreateTransaction", mock.AnythingOfType("domains.Transaction")).Return(domains.Transaction{}, nil)

	svc := NewTransactionService(repoMock, accSvc, ttSvc)

	_, err := svc.CreateTransaction(request.NewTransactionRequest{
		AccountId:       accId.String(),
		OperationTypeId: 1,
		Amount:          123.4,
	})

	assert.ErrorContains(t, err, "some error")
}
