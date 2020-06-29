package controller

import (
	"github.com/gin-gonic/gin"
	"gotest/model"
	"gotest/mysql"
	"net/http"
)


func Html(c *gin.Context)  {
	var result []model.UserModel
	mysql.Db.Select("email,password").Find(&result)

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"获取成功!",
		"data":result,
	})

}
