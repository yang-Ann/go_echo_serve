package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

var (
	// 服务端口
	PORT *int

	// 服务URL
	SERVE_URL string

	// 返回 body 数据的长度
	BODY_DATA_LENGTH = 100
)

func init() {
	// 可以通过命令行参数控制服务的端口
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


// 解码url
func QueryUnescape(str string) string {
	decodedStr, err := url.QueryUnescape(str)
	if err != nil {
		// fmt.Println("query 解码失败: ", err)
		// 如果错误则返回原始的字符串
		return str
	} else {
		return decodedStr
	}
}