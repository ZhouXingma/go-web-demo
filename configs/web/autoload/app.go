package autoload

// 程序相关的基本配置
type AppConfig struct {
	// 端口
	Port string `mapstructure:"port"`
}

var App = AppConfig{
	Port: "8080",
}
