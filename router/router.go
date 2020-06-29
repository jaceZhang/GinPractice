package router

import (
	"github.com/gin-gonic/gin"
	"gotest/controller"
	"gotest/middleware"
)

func RouterInit() *gin.Engine{
	r := gin.Default()
	r.Use(middleware.Cross())//跨域处理
	r.Static("/static","./static")
	r.StaticFile("/favicon.ico","./go.png")
	if mode := gin.Mode(); mode == gin.TestMode {
		r.LoadHTMLGlob("./../templates/*")
	} else {
		r.LoadHTMLGlob("templates/*")
	}

	g := r.Group("/")
	{
		g.Any("",controller.Index)
		g.POST("/register",controller.Register)
		g.POST("/login",controller.Login)
	}

	t := r.Group("/html")
	{
		t.Any("",controller.Html)
	}

	u:=r.Group("/user")
	{
		u.POST("/register",controller.Register)
		u.GET("/count",controller.GetCount)
		u.GET("/del_user",controller.DelUser)
	}

	return r
}
