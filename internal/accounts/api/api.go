package api

import (
	"github.com/gin-gonic/gin"
	accountsHandler "github.com/mateusffaria/pismo-challenge/internal/accounts/handlers"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/repositories"
	"github.com/mateusffaria/pismo-challenge/internal/accounts/services"
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
