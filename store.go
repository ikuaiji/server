package app

import (
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model

	Name string
}
