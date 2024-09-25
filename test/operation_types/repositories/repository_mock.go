package repositories_test

import (
	"pismo-challenge/internal/operation_types/domains"

	"github.com/stretchr/testify/mock"
)

type OperationTypesRepository struct {
	mock.Mock
}

func NewOperationTypesRepository() *OperationTypesRepository {
	return &OperationTypesRepository{}
}

func (m *OperationTypesRepository) GetOperationTypeById(id int) (domains.OperationType, error) {
	args := m.Called(id)

	return args.Get(0).(domains.OperationType), args.Error(1)
}
