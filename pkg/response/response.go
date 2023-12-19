package response

import (
	"github.com/gin-gonic/gin"
	"go-web-demo/pkg/httpcode"
)

// 返回json信息
func ReturnJson(c *gin.Context, httpCode httpcode.HttpCode, message *string, data interface{}) {
	messageInfo := httpCode.GetDesc()
	if nil != message {
		messageInfo = *message
	}
	c.JSONP(httpCode.GetCode(), gin.H{
		"code":    httpCode.GetCode(),
		"desc":    httpCode.GetDesc(),
		"success": httpCode.GetIsSuccess(),
		"message": messageInfo,
		"data":    data,
	})
}

// 成功信息
// 参数：
// c: 请求信息
// data：返回结果信息
func Success(c *gin.Context, data interface{}) {
	ReturnJson(c, httpcode.SUCCESS, nil, data)
}

// 失败信息
// 参数：
// c: 请求信息
// message：异常信息
func Fail(c *gin.Context, message string) {
	ReturnJson(c, httpcode.FAIL, &message, nil)
}
