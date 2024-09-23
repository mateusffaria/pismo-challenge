package domains

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID             uuid.UUID
	DocumentNumber string
}
