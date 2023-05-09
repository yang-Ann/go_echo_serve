package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Any("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/echo")
	})

	r.Any("/echo", handleEcho)

	r.Run(":8100")
}

func handleEcho(c *gin.Context) {
	fmt.Printf("echo %#v\n", c.Request)

	echoData := make(map[string]string)
	echoData["url"] = fmt.Sprintf("%s%s", c.Request.Host, c.Request.URL.String())
	echoData["method"] = c.Request.Method
	echoData["query"] = c.Request.URL.RawQuery
	for k, v := range c.Request.Header {
		echoData[k] = v[0]
	}
	echoData["remoteAddr"] = c.Request.RemoteAddr
	echoData["proto"] = c.Request.Proto
	fmt.Printf("PostForm: %#v\n", c.Request.PostForm)

	files, err := c.MultipartForm()
	if err == nil {
		for k, v := range files.Value {
			echoData["formData."+k] = strings.Join(v, "")
		}
	}

	// for _, v := range files.File["files"] {
	// 	echoData["formData." + v.Filename] = v.Filename;
	// }

	bytes, err := c.GetRawData()
	if err == nil {
		if len(bytes) < 500 {
			echoData["body"] = string(bytes)
		} else {
			echoData["body"] = string(bytes[:500]) + "省略其他..."
		}
	}

	c.JSON(http.StatusOK, echoData)
}
