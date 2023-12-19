package web

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	autoload2 "go-web-demo/configs/web/autoload"
	"strings"
	"sync"
)

// 配置信息整体未知
type Conf struct {
	// 环境信息
	Env string `mapstructure:"env"`
	// 系统基本配置信息
	AppConfig autoload2.AppConfig `mapstructure:"app"`
	// 日志配置信息
	LogConfig autoload2.LogConfig `mapstructure:"log"`
	// 数据库配置信息
	MysqlConfig autoload2.MysqlConfig `mapstructure:"mysql"`
}

var (
	// 默认配置信息
	Config = &Conf{
		// 环境信息
		Env:         "",
		AppConfig:   autoload2.App,
		LogConfig:   autoload2.Log,
		MysqlConfig: autoload2.Mysql,
	}
	// 用于执行一次配置初始化
	onece sync.Once
	// 用户配置信息的装噢诶
	V *viper.Viper
)

// 初始化配置文件
func InitConfig() {
	onece.Do(func() {
		// 加载配置文件
		load()
	})
}

// 加载配置文件
func load() {
	V = viper.New()
	// 加载公共配置
	loadOfCommon()
	// 加载命令行参数
	loadOfFlag()
	// 加载环境配置
	loadOfEvn()
}

// 加载公共配置
func loadOfCommon() {
	commConfigFilePath := "./configs/web/config.yml"
	V.SetConfigFile(commConfigFilePath)
	if err := V.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			panic("未找到公共配置: " + err.Error())
		} else {
			// 配置文件被找到，但产生了另外的错误
			panic("读取配置公共出错: " + err.Error())
		}
	}
	if err := V.Unmarshal(&Config); err != nil {
		panic(err)
	}
}

// 加载不同环境的环境配置
func loadOfEvn() {
	envConfigFilePath := strings.Builder{}
	envConfigFilePath.WriteString("./configs/web/config-")
	envConfigFilePath.WriteString(Config.Env)
	envConfigFilePath.WriteString(".yml")
	V.SetConfigFile(envConfigFilePath.String())
	if err := V.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			panic("未找到环境配置: " + err.Error())
		} else {
			// 配置文件被找到，但产生了另外的错误
			panic("读取配置环境出错: " + err.Error())
		}
	}
	if err := V.Unmarshal(&Config); err != nil {
		panic(err)
	}

}

func loadOfFlag() {
	var env string
	flag.StringVar(&env, "env", Config.Env, "环境变量")
	flag.Parse()
	fmt.Println("运行环境：", env)
	Config.Env = env
}
