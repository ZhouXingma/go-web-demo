package userapi

import (
	"github.com/gin-gonic/gin"
)

// 用户模块相关的路由信息
func UserRouter(e *gin.Engine) {
	g := e.Group("/user")
	{
		g.POST("/get", GetUser)
		g.POST("/add", AddUser)
		g.POST("/list", ListUser)
		g.POST("/remove", RemoveUser)
		g.POST("/page", PageUser)
		g.POST("/update", UpdateUser)
	}
}
