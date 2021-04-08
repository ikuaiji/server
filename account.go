package app

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model

	Name     string
	ParentId int
}
