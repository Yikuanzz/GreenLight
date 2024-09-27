package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// version 这里暂时将版本号用全局常量存储
const version = "1.0.0"

// config 包括服务器端口号和当前操作环境
type config struct {
	port int
	env  string
}

// application 定义处理函数、中间件的依赖
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// flag 解析命令存放到结构体中
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.Parse()

	// logger 创建日志将信息输出的终端
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// app 声明应用
	app := &application{
		config: cfg,
		logger: logger,
	}

	// mux 创建 HTTP 服务器并添加路由处理
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// srv 实例化一个服务器监听端口并提供时延设置
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// 启动服务器
	logger.Printf("Starting %s server on port %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
