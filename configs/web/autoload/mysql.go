package autoload

import (
	"go-web-demo/pkg/mysql"
	"unsafe"
)

// 数据库配置信息
type MysqlConfig struct {
	// mysql的url
	Url string `mapstructure:"url"`
	// 接口
	Port int `mapstructure:"port"`
	// 用户
	User string `mapstructure:"user"`
	// 密码
	Password string `mapstructure:"password"`
	// 数据库
	DbName string `mapstructure:"dbName"`
	// 其它配置
	OtherConfig string `mapstructure:"otherConfig"`
	// 最大连接数
	MaxOpenConns int `mapstructure:"maxOpenConns"`
	// 空闲连接池中的最大连接数
	MaxIdleConns int `mapstructure:"maxIdleConns"`
}

// 默认数据库配置信息
var Mysql = MysqlConfig{}

// 转换城mysql的配置
func (config MysqlConfig) Trans2MysqlConfig() *mysql.MysqlConfig {
	mysqlConfig := (*mysql.MysqlConfig)(unsafe.Pointer(&config))
	return mysqlConfig
}
