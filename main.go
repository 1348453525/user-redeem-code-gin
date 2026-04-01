package main

import (
	"net/http"

	"github.com/1348453525/user-redeem-code-gin/initialize"
)

func main() {
	// 初始化 Gin 引擎和路由
	r := initialize.InitRouter()
	// r.Run(":8080")

	// 配置 HTTP 服务
	src := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// 启动服务
	if err := src.ListenAndServe(); err != nil {
		panic(err)
	}
}
