package usecases_test

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	accountsAPI "pismo-challenge/internal/accounts/api"
	"pismo-challenge/internal/accounts/handlers/request"
	aResponse "pismo-challenge/internal/accounts/handlers/response"
	transactionsAPI "pismo-challenge/internal/transactions/api"
	tRequest "pismo-challenge/internal/transactions/handlers/request"
	tResponse "pismo-challenge/internal/transactions/handlers/response"

	dbSetup "pismo-challenge/test/integration/db"
)

func TestShouldCreateAccountTransactionSuccessfully(t *testing.T) {
	r := gin.Default()

	db := dbSetup.Init()

	accountsAPI.SetupApi(r, db)
	transactionsAPI.SetupApi(r, db)

	recorder := httptest.NewRecorder()

	account := request.AccountRequest{
		DocumentNumber: "1234567890",
	}
	accJson, _ := json.Marshal(account)

	r.ServeHTTP(recorder, httptest.NewRequest("POST", "/api/v1/accounts", strings.NewReader(string(accJson))))
	assert.Equal(t, 201, recorder.Code)

	var accountSaved aResponse.AccountResponse
	json.Unmarshal(recorder.Body.Bytes(), &accountSaved)

	transaction := tRequest.NewTransactionRequest{
		AccountId:       accountSaved.AccountId,
		OperationTypeId: 1,
		Amount:          123.45,
	}
	transactionJson, _ := json.Marshal(transaction)

	recorder = httptest.NewRecorder()

	r.ServeHTTP(recorder, httptest.NewRequest("POST", "/api/v1/transactions", strings.NewReader(string(transactionJson))))
	assert.Equal(t, 201, recorder.Code)

	var transactionSaved tResponse.NewTransactionResponse
	json.Unmarshal(recorder.Body.Bytes(), &transactionSaved)

	assert.Equal(t, transaction.AccountId, accountSaved.AccountId)
	assert.Equal(t, transaction.OperationTypeId, transactionSaved.OperationTypeId)
	assert.Equal(t, transaction.Amount, transactionSaved.Amount)

	dbSetup.Sunset(db)
}
