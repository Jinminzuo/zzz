package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

func main() {
	// 创建 Echo 实例
	e := echo.New()

	// 注册 /helloworld 路由
	e.GET("/helloworld", HelloWorldHandler)

	// 启动服务，监听 8080
	e.Logger.Fatal(e.Start(":8080"))
}

// HelloWorldHandler 返回 hello world JSON
func HelloWorldHandler(c echo.Context) error {
	result := JsonResult{
		Code: 0,
		Data: "helloworld echo",
	}
	return c.JSON(http.StatusOK, result)
}
