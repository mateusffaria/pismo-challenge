package repositories

import (
	"github.com/mateusffaria/pismo-challenge/internal/operation_types/domains"
	"gorm.io/gorm"
)

type OperationTypeRepositoryProvider interface {
	GetOperationTypeById(id int) (domains.OperationType, error)
}

type OperationType struct {
	DB *gorm.DB
}

func NewOperationType(db *gorm.DB) *OperationType {
	return &OperationType{
		DB: db,
	}
}

func (ar *OperationType) GetOperationTypeById(id int) (domains.OperationType, error) {
	var tt domains.OperationType

	res := ar.DB.First(&tt, "id = ?", id)
	if res.Error != nil {
		return tt, res.Error
	}

	return tt, nil
}
