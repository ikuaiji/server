package main

import (
	"app"
	"app/db"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouter.GET("/accounts", AccountsIndexHandler)
	apiRouter.GET("/account/:id", AccountGetHandler)
	apiRouter.POST("/account/:id", AccountPostHandler)
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

//AccountGetHandler 是GET /account/:id 接口的处理函数
func AccountGetHandler(c *gin.Context) {
	var param struct {
		ID uint `uri:"id"`
	}
	if err := c.ShouldBindUri(&param); err != nil {
		RenderError(c, err)
		return
	}

	record, err := db.GetAccount(param.ID)
	if err != nil {
		RenderError(c, err)
		return
	}

	RenderData(c, record)
}

//AccountPostHandler 是POST /account/:id 接口的处理函数
func AccountPostHandler(c *gin.Context) {
	var record app.Account
	if err := c.ShouldBindJSON(&record); err != nil {
		RenderError(c, err)
		return
	}

	var param struct {
		ID uint `uri:"id"`
	}
	if err := c.ShouldBindUri(&param); err != nil {
		RenderError(c, err)
		return
	}

	//始终以URI中的ID为准
	record.ID = param.ID

	err := db.Save(&record)
	if err != nil {
		RenderError(c, err)
		return
	}

	RenderData(c, true)
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
