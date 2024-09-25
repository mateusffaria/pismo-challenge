package api

import (
	accRepo "pismo-challenge/internal/accounts/repositories"
	accSvc "pismo-challenge/internal/accounts/services"
	ttRepo "pismo-challenge/internal/operation_types/repositories"
	ttSvc "pismo-challenge/internal/operation_types/services"
	accountsHandler "pismo-challenge/internal/transactions/handlers"
	"pismo-challenge/internal/transactions/repositories"
	"pismo-challenge/internal/transactions/services"

	"github.com/gin-gonic/gin"
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
