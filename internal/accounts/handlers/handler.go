package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/response"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/services"
	svcErrors "github.com/mateusffaria/pismo-challenge/internal/accounts/services/errors"
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
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})

		return
	}

	err = body.Validate()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})

		return
	}

	res, err := ah.asp.CreateUserAccount(body)
	if err != nil {
		switch {
		case errors.Is(err, svcErrors.ErrAccountDuplicated):
			c.JSON(http.StatusConflict, gin.H{"errors": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, response.UserAccountResponse{
		AccountId:      res.ID.String(),
		DocumentNumber: res.DocumentNumber,
	})
}

// Create User Account		godoc
// @Summary 							Get user's account
// @Description						Get the user account stored in DB
// @Param									id path string true "get account by id"
// @Produce 							application/json
// @Tags 									accounts
// @Success 							200 {object} response.UserAccountResponse
// @Router 								/v1/accounts/{id} [get]
func (ah AccountsHandler) GetUserAccount(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "missing identifier"})
		return
	}

	ua, err := ah.asp.GetUserAccount(id)
	if err != nil {
		switch {
		case errors.Is(err, svcErrors.ErrAccountNotFound):
			c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, response.UserAccountResponse{
		AccountId:      ua.ID.String(),
		DocumentNumber: ua.DocumentNumber,
	})
}
