package app

import (
	"time"

	"gorm.io/gorm"
)

//Bill 代表一次交易账单
type Bill struct {
	gorm.Model

	Type       string    //交易类型
	CategoryId uint      //科目ID
	AccountId  uint      //交易账户的ID
	Account2Id uint      //转账交易的对端账户
	Amount     float32   //金额
	BillAt     time.Time //入账时间
	MemberId   uint      //成员ID
	ProjectId  uint      //项目ID
	StoreId    uint      //商户ID
	Note       string    //交易备注
}
