package app

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	Name     string
	ParentId uint
}
