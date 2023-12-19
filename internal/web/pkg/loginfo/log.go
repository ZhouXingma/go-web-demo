package loginfo

import (
	"go-web-demo/pkg/json"
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

// 打印Info日志信息,主要用于方便json转换
// 参数：
// message: 前缀消息
// datas: 需要暂时的数据信息，比如参数信息
func Info(message string, datas ...interface{}) {
	if len(datas) > 0 {
		logInfos := make([]string, 1)
		logInfos = append(logInfos, message)
		for _, dataItem := range datas {
			jsonInfo, e := json.ToJsonString(dataItem)
			if e == nil {
				logInfos = append(logInfos, jsonInfo)
			} else {
				logInfos = append(logInfos, "[解析json异常："+e.Error()+"]")
			}
		}
		Log.Info(logInfos)
	} else {
		Log.Info(message)
	}
}

// 打印Error日志信息,主要用于方便json转换
// 参数：
// message: 前缀消息
// err: 异常信息
// datas: 需要暂时的数据信息，比如参数信息
func Error(message string, err error, datas ...interface{}) {
	if len(datas) > 0 {
		logInfos := make([]string, 2)
		logInfos = append(logInfos, message)
		logInfos = append(logInfos, err.Error())
		for _, dataItem := range datas {
			jsonInfo, e := json.ToJsonString(dataItem)
			if e != nil {
				logInfos = append(logInfos, jsonInfo)
			} else {
				logInfos = append(logInfos, "[解析json异常："+e.Error()+"]")
			}
		}
		Log.Error(logInfos)
	} else {
		Log.Error(message, err.Error())
	}
}
