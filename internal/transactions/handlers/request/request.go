package request

import (
	"fmt"

	validate "github.com/go-playground/validator/v10"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request/errors"
)

type NewTransactionRequest struct {
	AccountId       string  `json:"account_id,omitempty" validate:"required"`
	OperationTypeId int     `json:"operation_type_id,omitempty" validate:"required"`
	Amount          float32 `json:"amount,omitempty" validate:"required"`
}

func (uar *NewTransactionRequest) Validate() error {
	validator := validate.New(validate.WithRequiredStructEnabled())
	err := validator.Struct(uar)
	if err != nil {
		var fields = make([]string, 0)
		for _, err := range err.(validate.ValidationErrors) {
			fields = append(fields, fmt.Sprintf("%s is %s", err.StructField(), err.ActualTag()))
		}

		return errors.NewBodyError(fields)
	}

	return nil
}
