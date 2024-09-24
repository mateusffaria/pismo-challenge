package request

import (
	"fmt"

	validate "github.com/go-playground/validator/v10"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request/errors"
)

type AccountRequest struct {
	DocumentNumber string `json:"document_number,omitempty" validate:"required"`
}

func (uar *AccountRequest) Validate() error {
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
