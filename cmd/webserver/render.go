package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//错误码定义
const (
	ErrNormal  = iota //无错误
	ErrGeneral        //通用错误码（无明确定义）
)

//RenderErrorCode 输出带错误码带错误信息
func RenderErrorCode(c *gin.Context, code int, err error) {
	c.JSON(http.StatusOK, gin.H{"code": code, "message": err.Error()})
}

//RenderError 输出通用错误码的错误信息
func RenderError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"code": ErrGeneral, "message": err.Error()})
}

//RenderData 输出正常的业务数据
func RenderData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": ErrNormal, "data": data})
}
