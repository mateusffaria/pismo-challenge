package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/response"
	customError "github.com/mateusffaria/pismo-challenge/internal/accounts/repositories/errors"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/services"
)

type AccountsHandler struct {
	asp services.AccountServiceProvider
}

func NewAccountsHandler(asp services.AccountServiceProvider) *AccountsHandler {
	return &AccountsHandler{
		asp: asp,
	}
}

// Create User Account		godoc
// @Summary 							Create a new user account
// @Description						Save a new user in DB
// @Param									account body request.UserAccountRequest true "Create user account"
// @Produce 							application/json
// @Tags 									accounts
// @Success 							201 {object} response.UserAccountResponse
// @Router 								/v1/accounts [post]
func (ah AccountsHandler) CreateUserAccount(c *gin.Context) {
	body := request.UserAccountRequest{}

	err := c.ShouldBindBodyWithJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	err = body.Validate()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": err.Error(),
		})

		return
	}

	res, err := ah.asp.CreateUserAccount(body)
	if errors.Is(err, &customError.DuplicateEntity{}) {
		c.JSON(http.StatusConflict, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.UserAccountResponse{
		AccountId:      res.ID.String(),
		DocumentNumber: res.DocumentNumber,
	})
}
