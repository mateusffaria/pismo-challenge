package services

import (
	"errors"

	"github.com/mateusffaria/pismo-challenge/internal/operation_types/domains"
	"github.com/mateusffaria/pismo-challenge/internal/operation_types/repositories"
	customErrors "github.com/mateusffaria/pismo-challenge/internal/operation_types/services/errors"
	"gorm.io/gorm"
)

type OperationTypesServiceProvider interface {
	GetOperationType(OperationTypeId int) (domains.OperationType, error)
}

type OperationTypesService struct {
	tr repositories.OperationTypeRepositoryProvider
}

func NewOperationTypesService(tr repositories.OperationTypeRepositoryProvider) *OperationTypesService {
	return &OperationTypesService{
		tr: tr,
	}
}

func (ts *OperationTypesService) GetOperationType(OperationTypeId int) (domains.OperationType, error) {
	tt, err := ts.tr.GetOperationTypeById(OperationTypeId)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return tt, customErrors.ErrNotFound
		default:
			return tt, customErrors.ErrInternalDatabaseError
		}
	}

	return tt, nil
}
