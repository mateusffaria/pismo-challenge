package domains

import (
	"gorm.io/gorm"
)

type TransactionType struct {
	gorm.Model
	Description string
}
