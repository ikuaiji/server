package main

import (
	"app"
	"app/db"
	"flag"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var dsn, listen string

func init() {
	flag.StringVar(&dsn, "dsn", "root:root@tcp(127.0.0.1:3306)/laokuaiji", "数据库DSN")
	flag.StringVar(&listen, "listen", ":8081", "侦听地址")
}

func main() {
	flag.Parse()

	//初始化数据库
	err := db.Init(dsn)
	if err != nil {
		log.Fatal(err)
	}

	//初始化Gin并注册路由
	r := gin.Default()
	r.GET("/bills", BillsIndexHandler)

	//启动HTTP侦听器
	r.Run(listen)
}

//BillsIndexHandler 是GET /bills接口的处理函数
func BillsIndexHandler(c *gin.Context) {
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

	bills, err := db.GetBillsOfMonth(param.Year, param.Month)
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