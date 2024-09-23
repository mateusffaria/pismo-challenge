package domains

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID             string
	DocumentNumber string
}
