package handlers

import (
	"errors"
	"net/http"

	"pismo-challenge/internal/accounts/handlers/request"
	"pismo-challenge/internal/accounts/handlers/response"
	"pismo-challenge/internal/accounts/services"
	svcErrors "pismo-challenge/internal/accounts/services/errors"

	"github.com/gin-gonic/gin"
)

type AccountsHandler struct {
	asp services.AccountServiceProvider
}

func NewAccountsHandler(asp services.AccountServiceProvider) *AccountsHandler {
	return &AccountsHandler{
		asp: asp,
	}
}

// Create  Account		godoc
// @Summary 							Create a new  account
// @Description						Save a new  in DB
// @Param									account body request.AccountRequest true "Create  account"
// @Produce 							application/json
// @Tags 									accounts
// @Success 							201 {object} response.AccountResponse
// @Router 								/v1/accounts [post]
func (ah AccountsHandler) CreateAccount(c *gin.Context) {
	body := request.AccountRequest{}

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

	res, err := ah.asp.CreateAccount(body)
	if err != nil {
		switch {
		case errors.Is(err, svcErrors.ErrAccountDuplicated):
			c.JSON(http.StatusConflict, gin.H{"errors": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, response.AccountResponse{
		AccountId:      res.ID.String(),
		DocumentNumber: res.DocumentNumber,
	})
}

// Create  Account		godoc
// @Summary 							Get 's account
// @Description						Get the  account stored in DB
// @Param									id path string true "get account by id"
// @Produce 							application/json
// @Tags 									accounts
// @Success 							200 {object} response.AccountResponse
// @Router 								/v1/accounts/{id} [get]
func (ah AccountsHandler) GetAccount(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "missing identifier"})
		return
	}

	ua, err := ah.asp.GetAccount(id)
	if err != nil {
		switch {
		case errors.Is(err, svcErrors.ErrAccountNotFound):
			c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, response.AccountResponse{
		AccountId:      ua.ID.String(),
		DocumentNumber: ua.DocumentNumber,
	})
}
