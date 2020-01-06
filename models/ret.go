package models

import (
	"encoding/json"
	"net/http"
)

//HTTPRet, 为了配合go-kit默认的错误处理函数
//需要实现以下几种接口：json.Marshaler,transport的StatusCoder，error
type Ret struct {
	Code         int
	Msg          string
	TokenInvalid bool //如果该值为true，前端需要重新登陆
	Data         interface{}
}

//error
func (h Ret) Error() string {
	return h.Msg
}

//statusCoder
func (h Ret) StatusCode() int {
	return http.StatusOK //???为什么要固定为statusOK
}

//json.Marshaler
func (h Ret) MarshalJson() ([]byte, error) {
	//克隆是为了防止在调用Json.Marshal时发生以下错误	>>>???不太明白
	//runtime: goroutine stack exceeds 100000000-byte limit
	//fatal error: stack overflow
	clone := struct {
		Code         int         `json:"code"`
		Msg          string      `json:"msg"`
		TokenInvalid bool        `json:"tokenInValid"`
		Data         interface{} `json:"data"`
	}{
		Code:         h.Code,
		Msg:          h.Msg,
		TokenInvalid: h.TokenInvalid,
		Data:         h.Data,
	}
	return json.Marshal(&clone)
}
