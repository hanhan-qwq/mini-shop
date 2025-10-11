package main

import (
	"fmt"
	"mini_shop/config"
	"mini_shop/global"
	"mini_shop/web/router"
	"net/http"
	"os"
)

func main() {
	config.InitConfig("yaml/config.yaml")
	global.InitMysql()

	r := router.InitRouter()
	addr := fmt.Sprintf("%s:%d", "localhost", config.AppConfig.Server.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("服务器启动失败", err)
		os.Exit(-1)
	}
}
