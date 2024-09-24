package domains

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID              uuid.UUID
	AccountId       uuid.UUID
	OperationTypeId int
	Amount          float32
	EventDate       time.Time
}
