package main

import (
	"app/db"
	"flag"
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

var dsn, listen, staticPath string
var ginEngine = gin.Default()
var apiRouter = ginEngine.Group("/api")

func init() {
	flag.StringVar(&dsn, "dsn", "root:root@tcp(127.0.0.1:3306)/laokuaiji", "数据库DSN")
	flag.StringVar(&listen, "listen", ":8081", "侦听地址")
	flag.StringVar(&staticPath, "static", "../web/dist", "前端静态文件的路径")
}

func main() {
	flag.Parse()

	//初始化数据库
	err := db.Init(dsn)
	if err != nil {
		log.Fatal(err)
	}

	ginEngine.NoRoute(StaticFileHandler)

	//启动HTTP侦听器
	ginEngine.Run(listen)
}

//StaticFileHandler 用于处理前端提供的静态文件
func StaticFileHandler(c *gin.Context) {
	//判断请求路径是否是文件，如果是则返回
	file := path.Join(staticPath, c.Request.URL.Path)
	if fi, _ := os.Stat(file); fi != nil && fi.Mode().IsRegular() {
		c.File(file)
		return
	}

	file = path.Join(file, "index.html")
	//如果不是，则尝试改路径下的index.html文件，如果是则返回
	//如果不是，按缺省设置处理（是文件夹则返回文件列表、其他返回404）
	if fi, _ := os.Stat(file); fi != nil && fi.Mode().IsRegular() {
		c.File(file)
	}
}
