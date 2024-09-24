package domains

import (
	"gorm.io/gorm"
)

type OperationType struct {
	gorm.Model
	Description string
}
