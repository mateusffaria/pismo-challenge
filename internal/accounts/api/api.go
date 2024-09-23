package api

import (
	"github.com/gin-gonic/gin"
	accountsHandler "github.com/mateusffaria/pismo-challenge/internal/accounts/handlers"
)

func SetupApi(r *gin.Engine) {
	ah := accountsHandler.NewAccountsHandler()

	group := r.Group("/api/v1")
	{
		group.POST("/accounts", ah.CreateUserAccount)
	}
}
