package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

// 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 请求前
		c.Next()
		// 请求后

		// 计算处理时间
		latency := time.Since(t)

		url := c.Request.URL
		method := c.Request.Method

		remoteAddr := c.Request.RemoteAddr
		// proto := c.Request.Proto

		// 获取发送的 status
		status := c.Writer.Status()

		decodedStr := QueryUnescape(url.String())
		color.New(color.FgHiGreen).Print(decodedStr)
		fmt.Printf("\t")
		color.New(color.FgYellow, color.Bold).Print(method)
		fmt.Printf("\t%v", remoteAddr)
		fmt.Printf("\t %v", latency)
		fmt.Printf("\t")

		switch status {
		case http.StatusOK:
			color.New(color.FgGreen, color.Bold).Print(status)
		default:
			color.New(color.FgRed, color.Bold).Print(status)
		}

		fmt.Println()
	}
}
