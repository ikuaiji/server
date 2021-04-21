package main

import (
	"app/db"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

var dsn, listen string
var ginEngine = gin.Default()
var apiRouter = ginEngine.Group("/api")

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

	//启动HTTP侦听器
	ginEngine.Run(listen)
}
