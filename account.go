package app

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model

	Type     string //账户类型: bank银行账户、virtual虚拟账户、debt借贷账户
	Name     string
	Hide     bool //账户是否隐藏显示
	ParentId int
}
