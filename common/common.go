package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//json函数封装
func OutJson(c *gin.Context,code int,msg string,data interface{})  {
	c.JSON(http.StatusOK,gin.H{
		"code" : code,
		"msg" : msg,
		"data" : data,
	})
}
