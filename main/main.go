package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

var (
	PORT      *int
	SERVE_URL string
)

func init() {
	PORT = flag.Int("port", 8100, "端口号")
	flag.Parse()

	SERVE_URL = fmt.Sprintf("http://localhost:%v", *PORT)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	// r := gin.Default()

	r := gin.New()
	r.Use(gin.Recovery())
	// 使用自定义日志中间件
	r.Use(Logger())

	// 任意请求
	r.Any("/*action", handleEcho)

	LogTip()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", *PORT),
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// 打印提示
func LogTip() {
	fmt.Println()
	color.New(color.FgGreen, color.Bold, color.Underline).Print(SERVE_URL)
	fmt.Print(" 服务启动成功")
	fmt.Printf("\n\n")
}
