package services_test

import (
	"github.com/mateusffaria/pismo-challenge/internal/operation_types/domains"
	"github.com/stretchr/testify/mock"
)

type OperationTypesServiceRepository struct {
	mock.Mock
}

func NewOperationTypesService() *OperationTypesServiceRepository {
	return &OperationTypesServiceRepository{}
}

func (m *OperationTypesServiceRepository) GetOperationType(OperationTypeId int) (domains.OperationType, error) {
	args := m.Called(OperationTypeId)

	return args.Get(0).(domains.OperationType), args.Error(1)
}
