package model

// ResponseBody 默认的返回类型
type ResponseBody struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	E    interface{} `json:"e"`
}
