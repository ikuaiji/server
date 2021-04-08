package app

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model

	Name string
}
