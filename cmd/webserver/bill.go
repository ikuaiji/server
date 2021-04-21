package main

import (
	"app"
	"app/db"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouter.GET("/bills", BillsIndexHandler)
	apiRouter.GET("/bill/:id", BillGetHandler)
	apiRouter.POST("/bill/:id", BillPostHandler)
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

func BillGetHandler(c *gin.Context) {
	var param struct {
		ID uint `uri:"id"`
	}
	if err := c.ShouldBindUri(&param); err != nil {
		RenderError(c, err)
		return
	}

	bill, err := db.GetBill(param.ID)
	if err != nil {
		RenderError(c, err)
		return
	}

	idNames, err := db.GetMetaIdNameMap()
	if err != nil {
		RenderError(c, err)
		return
	}

	RenderData(c, gin.H{"bill": bill, "id_names": idNames})
}

func BillPostHandler(c *gin.Context) {
	var bill app.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
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
	bill.ID = param.ID

	err := db.Save(&bill)
	if err != nil {
		RenderError(c, err)
		return
	}

	RenderData(c, true)
}
