package autoload

import (
	"go-web-demo/pkg/sglog"
	"unsafe"
)

// 日志配置
type LogConfig struct {
	// 日志文件
	File string `mapstructure:"file"`
	// 文件大小限制,单位MB
	MaxSize int `mapstructure:"maxSize"`
	// 最大保留日志文件数量
	MaxBackups int `mapstructure:"maxBackups"`
	// 日志文件保留天数
	MaxAge int `mapstructure:"maxAge"`
	// 是否压缩处理
	Compress bool `mapstructure:"compress"`
}

// 默认日志配置
var Log = LogConfig{
	File:       "/data/go-web-demo/go-web-demo.log",
	MaxSize:    100,
	MaxBackups: 200,
	MaxAge:     30,
	Compress:   true,
}

// 自定义的日志配置信息转换为日志初始化需要的日志配置信息
// 返回: 系统初始化需要的日志配置信息，参数比较齐全，主要用于日志初始化
func (logConfig LogConfig) Trans2SgLogConfig() *sglog.SgLogConfig {
	sgLogConfig := (*sglog.SgLogConfig)(unsafe.Pointer(&logConfig))
	return sgLogConfig
}
