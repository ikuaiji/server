package main

import (
	"app/db"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouter.GET("/summary", SummaryGetHandler)
}

//AccountsIndexHandler 是GET /accounts接口的处理函数
func SummaryGetHandler(c *gin.Context) {
	from, err := db.GetFirstBillAt()
	if err != nil {
		RenderError(c, err)
		return
	}

	count, err := db.GetBillsCount()
	if err != nil {
		RenderError(c, err)
		return
	}

	balance, err := db.GetBalance()
	if err != nil {
		RenderError(c, err)
		return
	}

	respData := gin.H{
		"From":         from,
		"RecordsCount": count,
		"Balance":      balance,
	}
	RenderData(c, respData)
}
