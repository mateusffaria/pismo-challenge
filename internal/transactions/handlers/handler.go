package handlers

import (
	"errors"
	"fmt"
	"net/http"

	asErrors "pismo-challenge/internal/accounts/services/errors"
	ttErrors "pismo-challenge/internal/operation_types/services/errors"
	"pismo-challenge/internal/transactions/handlers/request"
	"pismo-challenge/internal/transactions/handlers/response"
	"pismo-challenge/internal/transactions/services"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	tsp services.TransactionServiceProvider
}

func NewTransactionHandler(tsp services.TransactionServiceProvider) *TransactionHandler {
	return &TransactionHandler{
		tsp: tsp,
	}
}

// Create  Transaction		godoc
// @Summary 									Create a new  transaction
// @Description								Save a new  in DB
// @Param											transaction body request.NewTransactionRequest true "Create trasaction for a given "
// @Produce 									application/json
// @Tags 											transactions
// @Success 									201 {object} response.NewTransactionResponse
// @Router 										/v1/transactions [post]
func (th TransactionHandler) CreateTransaction(c *gin.Context) {
	body := request.NewTransactionRequest{}

	err := c.ShouldBindBodyWithJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})

		return
	}

	err = body.Validate()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})

		return
	}

	res, err := th.tsp.CreateTransaction(body)
	if err != nil {
		fmt.Printf("error %v", err)
		switch {
		case (errors.Is(err, ttErrors.ErrNotFound) || errors.Is(err, asErrors.ErrAccountNotFound)):
			c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		}

		return
	}

	fAmount, _ := res.Amount.Float64()

	c.JSON(http.StatusCreated, response.NewTransactionResponse{
		TransactionId:   res.ID.String(),
		AccountId:       res.ID.String(),
		OperationTypeId: res.OperationTypeId,
		Amount:          fAmount,
	})
}
