package sglog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type SgLogConfig struct {
	// 日志文件
	File string
	// 文件大小限制,单位MB
	MaxSize int
	// 最大保留日志文件数量
	MaxBackups int
	// 日志文件保留天数
	MaxAge int
	// 是否压缩处理
	Compress bool
}

// 初始化日志,在bootstarp中进行初始化
func InitLog(logConfig *SgLogConfig) *zap.SugaredLogger {
	//编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	//时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//日志等级字母大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logConfig.File,       //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    logConfig.MaxSize,    //文件大小限制,单位MB
		MaxBackups: logConfig.MaxBackups, //最大保留日志文件数量
		MaxAge:     logConfig.MaxAge,     //日志文件保留天数
		Compress:   logConfig.Compress,   //是否压缩处理
	})
	//第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer,
		zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
	//AddCaller()为显示文件名和行号
	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}
