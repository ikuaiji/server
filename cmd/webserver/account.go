package main

import (
	"app/db"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouter.GET("/accounts", AccountsIndexHandler)
	apiRouter.GET("/account_balances", AccountBalancesIndexHandler)
}

//AccountsIndexHandler 是GET /accounts接口的处理函数
func AccountsIndexHandler(c *gin.Context) {
	accounts, err := db.GetAccounts()
	if err != nil {
		RenderError(c, err)
		return
	}

	RenderData(c, accounts)
}

//AccountBalancesIndexHandler 是GET /account_balances接口的处理函数
func AccountBalancesIndexHandler(c *gin.Context) {
	balances, err := db.GetAccountBalances()
	if err != nil {
		RenderError(c, err)
		return
	}

	RenderData(c, balances)
}
