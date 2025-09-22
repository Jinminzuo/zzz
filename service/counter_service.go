package service

import (
	"encoding/json"
	"net/http"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// HelloWorldHandler 返回 hello world
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// 仅支持 GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(JsonResult{
			Code:     -1,
			ErrorMsg: "请求方法不支持",
		})
		return
	}

	// 构建返回结果
	result := JsonResult{
		Code: 0,
		Data: "helloworld",
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
