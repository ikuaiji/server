package app

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model

	Type       string
	CategoryId uint
	AccountId  uint
	Amount     float32
	BillAt     time.Time //入账时间
	MemberId   uint
	ProjectId  uint
	StoreId    uint
	Note       string
}
