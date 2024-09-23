package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/handlers/request"
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
// @Success 							201
// @Router 								/api/v1/accounts [post]
func (ah AccountsHandler) CreateUserAccount(c *gin.Context) {
	body := request.UserAccountRequest{}

	err := c.ShouldBindBodyWithJSON(body)
	if err != nil {
		log.Default().Println("error binding user body values")
		c.JSON(500, gin.H{
			"errors": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
