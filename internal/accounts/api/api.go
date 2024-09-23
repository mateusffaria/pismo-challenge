package api

import (
	"github.com/gin-gonic/gin"
	accountsHandler "github.com/mateusffaria/pismo-challenge/internal/accounts/handlers"
)

func SetupApi(r *gin.Engine) {
	ah := accountsHandler.NewAccountsHandler()

	r.Group("v1")
	{
		r.POST("/accounts", ah.CreateUserAccount)
	}
}
