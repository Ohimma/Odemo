package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ohimma/odemo/gin/config"
	"github.com/ohimma/odemo/gin/middleware"
	"github.com/ohimma/odemo/gin/model"
	"github.com/ohimma/odemo/gin/router"
)

func main() {
	fmt.Println("run server")

	// 1. 定义配置文件
	// fmt.Println("main = ", config.Conf.Mysql.Host)

	// 2. 初始化数据库
	model.InitMysql()

	if config.Conf.Server.UseRedis {
		model.InitRedis()
	}

	// 3. 初始化log
	// 4. 进入 gin 主程序
	runServer()
}

func runServer() {
	router1 := router.InitRouter()

	// 优雅关停
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.Server.Port),
		Handler: router1,
	}

	middleware.Logger.Info(fmt.Sprintf("Listening and serving HTTP on Port: %d, Pid: %d", config.Conf.Server.Port, os.Getpid()))

	server.ListenAndServe()
	middleware.Logger.Info("server exit .....")
}
