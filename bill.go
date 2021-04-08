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
	Account2Id uint      //转账交易的对端账户
	BillAt     time.Time //入账时间
	MemberId   uint
	ProjectId  uint
	StoreId    uint
	Note       string
}
