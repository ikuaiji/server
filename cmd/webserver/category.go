package main

import (
	"app"
	"app/db"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouter.GET("/categories", CategoriesIndexHandler)
	apiRouter.GET("/category_sums", CategorySumsIndexHandler)
}

//CategoriesIndexHandler 是GET /categories接口的处理函数
func CategoriesIndexHandler(c *gin.Context) {
	records, err := db.GetCategories()
	if err != nil {
		RenderError(c, err)
		return
	}

	RenderData(c, records)
}

//CategorySumsIndexHandler 是GET /category_sums接口的处理函数
func CategorySumsIndexHandler(c *gin.Context) {
	var param struct {
		Year  int        `form:"year"`
		Month time.Month `form:"month"`
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

	result, err := db.GetCategorySums(param.Year, param.Month)
	if err != nil {
		RenderError(c, err)
		return
	}

	RenderData(c, result)
}
