package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-web-demo/internal/web/pkg/loginfo"
)

type MysqlConfig struct {
	// mysql的url
	Url string
	// 接口
	Port int
	// 用户
	User string
	// 密码
	Password string
	// 数据库
	DbName string
	// 其它配置
	OtherConfig string
	// 最大连接数
	MaxOpenConns int
	// 空闲连接池中的最大连接数
	MaxIdleConns int
}

func (config MysqlConfig) IsValid() bool {
	return config.Url != "" && config.Port != 0 && config.User != "" && config.Password != "" && config.DbName != ""
}

// 初始化数据库
// 参数：
// config：配置信息
func InitDb(config *MysqlConfig) *sqlx.DB {
	if !config.IsValid() {
		return nil
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.User, config.Password, config.Url, config.Port, config.DbName, config.OtherConfig)
	db, err := sqlx.Connect("mysql", dsn)
	loginfo.Log.Info(dsn)
	if err != nil {
		loginfo.Log.Error("链接mysql数据库失败:%v\n", err)
		panic("链接mysql数据库失败")
	} else {
		loginfo.Log.Info("链接mysql数据库成功", nil)
	}
	// 设置数据库的最大打开连接数。
	db.SetMaxOpenConns(config.MaxOpenConns)
	// 设置空闲连接池中的最大连接数。
	db.SetMaxIdleConns(config.MaxIdleConns)
	return db
}
