package bootstrap

import (
	"go-web-demo/configs/web"
	"go-web-demo/internal/web/pkg/content"
	"go-web-demo/internal/web/pkg/loginfo"
	"go-web-demo/internal/web/routers"
	"go-web-demo/pkg/mysql"
	"go-web-demo/pkg/sglog"
)

func Start() {
	Init()
	Run()
}

// 初始化配置和数据
func Init() {
	// 初始化配置
	web.InitConfig()
	// 初始化日志
	InitLog()
	// 初始化数据库
	InitDb()
}

// 开始运行
func Run() {
	r := routers.InitRouter()
	r.Run(":" + web.Config.AppConfig.Port) // 监听并在 0.0.0.0:8080 上启动服务
}

// 初始化日志
func InitLog() {
	logConfig := web.Config.LogConfig.Trans2SgLogConfig()
	loginfo.Log = sglog.InitLog(logConfig)
}

// 初始化数据库
func InitDb() {
	dbConfig := web.Config.MysqlConfig.Trans2MysqlConfig()
	content.Db = mysql.InitDb(dbConfig)
}
