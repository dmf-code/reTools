package main

import (
	"reTools/web/app/routers"
)

//go:generate go version

func main() {

	r := routers.InitRouter()

	r.Run("0.0.0.0:28888") // 监听并在 0.0.0.0:8080 上启动服务
}
