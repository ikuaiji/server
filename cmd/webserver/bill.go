package main

import (
	"app"
	"app/db"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouter.GET("/bills", BillsGetHandler)
	apiRouter.POST("/bills", BillsPostHandler)

	apiRouter.GET("/bill/:id", BillGetHandler)
	apiRouter.POST("/bill/:id", BillPostHandler)
	apiRouter.DELETE("/bill/:id", BillDeleteHandler)
}

//BillsIndexHandler 是GET /bills接口的处理函数
func BillsGetHandler(c *gin.Context) {
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

//BillsPostHandler 是POST /bills接口的处理函数
func BillsPostHandler(c *gin.Context) {
	var bill app.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		RenderError(c, err)
		return
	}

	//始终以URI中的ID为准
	if bill.ID > 0 {
		RenderError(c, fmt.Errorf("Invalid bill attribute: id should be empty"))
		return
	}

	err := db.Save(&bill)
	if err != nil {
		RenderError(c, err)
		return
	}

	RenderData(c, bill)
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

//BillDeleteHandler 是DELETE /bill/:id 接口的处理函数
func BillDeleteHandler(c *gin.Context) {
	var param struct {
		ID uint `uri:"id"`
	}
	if err := c.ShouldBindUri(&param); err != nil {
		RenderError(c, err)
		return
	}

	if param.ID == 0 {
		RenderError(c, fmt.Errorf("Invalid bill id (%d) to delete", param.ID))
		return
	}

	var bill app.Bill
	bill.ID = param.ID

	err := db.Delete(&bill)
	if err != nil {
		RenderError(c, err)
		return
	}

	RenderData(c, true)
}
