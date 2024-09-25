package api

import (
	accountsHandler "pismo-challenge/internal/accounts/handlers"
	"pismo-challenge/internal/accounts/repositories"
	"pismo-challenge/internal/accounts/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupApi(r *gin.Engine, db *gorm.DB) {
	ar := repositories.NewAccountRepository(db)
	as := services.NewAccountService(ar)
	ah := accountsHandler.NewAccountsHandler(as)

	group := r.Group("/api/v1")
	{
		group.POST("/accounts", ah.CreateAccount)
		group.GET("/accounts/:id", ah.GetAccount)
	}
}
