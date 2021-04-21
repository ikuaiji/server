package main

import (
	"app"
	"app/db"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouter.GET("/bills", BillsIndexHandler)
}

//BillsIndexHandler 是GET /bills接口的处理函数
func BillsIndexHandler(c *gin.Context) {
	var param struct {
		Year      int        `form:"year"`
		Month     time.Month `form:"month"`
		AccountID uint       `form:"account_id"`
	}

	if err := c.ShouldBindQuery(&param); err != nil {
		RenderError(c, err)
		return
	}

	now := time.Now().In(app.TZ)
	if param.Year == 0 {
		param.Year = now.Year()
	}
	if param.Month == 0 {
		param.Month = now.Month()
	}

	bills, err := db.GetBillsOfMonth(param.Year, param.Month, param.AccountID)
	if err != nil {
		RenderError(c, err)
		return
	}

	idNames, err := db.GetMetaIdNameMap()
	if err != nil {
		RenderError(c, err)
		return
	}

	RenderData(c, gin.H{"bills": bills, "id_names": idNames})
}
