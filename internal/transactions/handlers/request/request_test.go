package request_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/handlers/request"
	"github.com/stretchr/testify/assert"
)

func TestShouldValidateAccountRequestSuccessfully(t *testing.T) {
	req := request.NewTransactionRequest{
		AccountId:       uuid.New().String(),
		OperationTypeId: 1,
		Amount:          123,
	}

	err := req.Validate()

	assert.NoError(t, err)
}

func TestShouldReturnErrorWhenInvalidBody(t *testing.T) {
	req := request.NewTransactionRequest{}

	err := req.Validate()

	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid body request")
}
