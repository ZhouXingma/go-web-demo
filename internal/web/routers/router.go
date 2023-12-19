package routers

import (
	"github.com/gin-gonic/gin"
	"go-web-demo/internal/web/api/userapi"
	"net/http"
)

var (
	engine *gin.Engine
)

// 初始化路由
func InitRouter() *gin.Engine {
	engine = gin.New()
	// 跨域问题处理
	engine.Use(cors())
	// 健康检测
	healthCheck(engine)
	// 其它路由地址
	addRouters(engine)
	return engine
}

// 健康检测
// 参数：
// e: 请求信息，gin框架请求的上下文信息
func healthCheck(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
}

// 添加其它路由地址
// 备注： 如果后续有新的模块，需要从这里加入配置信息
// 参数：
// e: 请求信息，gin框架请求的上下文信息
func addRouters(e *gin.Engine) {
	userapi.UserRouter(e)
}

// 添加跨域支持
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("origin") //origin值获取
		if origin != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
