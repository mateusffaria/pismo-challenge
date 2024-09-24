package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/handlers/request"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/handlers/response"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/services"
)

type TransactionHandler struct {
	asp services.TransactionServiceProvider
}

func NewTransactionHandler(asp services.TransactionServiceProvider) *TransactionHandler {
	return &TransactionHandler{
		asp: asp,
	}
}

// Create User Account		godoc
// @Summary 							Create a new user account
// @Description						Save a new user in DB
// @Param									account body request.NewTransactionRequest true "Create trasaction for a given user"
// @Produce 							application/json
// @Tags 									accounts
// @Success 							201 {object} response.NewTransactionResponse
// @Router 								/v1/transactions [post]
func (ah TransactionHandler) CreateTransaction(c *gin.Context) {
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

	res, err := ah.asp.CreateTransaction(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
	}

	// TODO: Adjust response
	c.JSON(http.StatusCreated, response.NewTransactionResponse{
		AccountId: res.ID.String(),
	})
}
