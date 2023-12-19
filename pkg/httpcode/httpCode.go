package httpcode

type HttpCode int

// API请求给前端的编码
const (
	SUCCESS HttpCode = 200
	FAIL    HttpCode = 500
)

// 获取编码
func (httpCode HttpCode) GetCode() int {
	return int(httpCode)
}

// 获取描述
func (httpCode HttpCode) GetDesc() string {
	code := int(httpCode)
	des := ""
	switch code {
	case 200:
		des = "成功"
	case 500:
		des = "失败"
	}
	return des
}

// 是否成功
func (httpCode HttpCode) GetIsSuccess() bool {
	code := int(httpCode)
	isSuccess := false
	switch code {
	case 200:
		isSuccess = true
	case 500:
		isSuccess = false
	}
	return isSuccess
}
