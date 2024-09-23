package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/response"
)

type AccountsHandler struct {
}

func NewAccountsHandler() *AccountsHandler {
	return &AccountsHandler{}
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
		fmt.Println("error binding user body values")
		c.JSON(500, gin.H{
			"errors": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response.UserAccountResponse{
		DocumentNumber: body.DocumentNumber,
	})
}
