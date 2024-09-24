package domains

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID              uuid.UUID
	AccountId       uuid.UUID
	OperationTypeId int
	Amount          decimal.Decimal
	EventDate       time.Time
}
