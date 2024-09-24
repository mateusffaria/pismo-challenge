package api

import (
	"github.com/gin-gonic/gin"
	accRepo "github.com/mateusffaria/pismo-challenge/internal/accounts/repositories"
	accSvc "github.com/mateusffaria/pismo-challenge/internal/accounts/services"
	ttRepo "github.com/mateusffaria/pismo-challenge/internal/transaction_types/repositories"
	ttSvc "github.com/mateusffaria/pismo-challenge/internal/transaction_types/services"
	accountsHandler "github.com/mateusffaria/pismo-challenge/internal/transactions/handlers"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/repositories"
	"github.com/mateusffaria/pismo-challenge/internal/transactions/services"
	"gorm.io/gorm"
)

func SetupApi(r *gin.Engine, db *gorm.DB) {
	ar := accRepo.NewAccountRepository(db)
	as := accSvc.NewAccountService(ar)

	ttr := ttRepo.NewOperationType(db)
	tts := ttSvc.NewOperationTypesService(ttr)

	tr := repositories.NewTransactionRepository(db)
	ts := services.NewTransactionService(tr, as, tts)
	th := accountsHandler.NewTransactionHandler(ts)

	group := r.Group("/api/v1")
	{
		group.POST("/transactions", th.CreateTransaction)
	}
}
