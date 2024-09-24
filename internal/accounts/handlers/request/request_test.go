package request_test

import (
	"testing"

	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request"
	"github.com/stretchr/testify/assert"
)

func TestShouldValidateAccountRequestSuccessfully(t *testing.T) {
	req := request.AccountRequest{
		DocumentNumber: "12345",
	}

	err := req.Validate()

	assert.NoError(t, err)
}

func TestShouldReturnErrorWhenInvalidBody(t *testing.T) {
	req := request.AccountRequest{}

	err := req.Validate()

	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid body request")
}
