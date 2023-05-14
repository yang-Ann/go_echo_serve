package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 响应数据结构
type EchoData struct {
	Url           string            `json:"url"`
	Method        string            `json:"method"`
	Header        map[string]string `json:"header"`
	Body          *string           `json:"body,omitempty"`
	Query         map[string]string `json:"query,omitempty"`
	FormDataValue map[string]string `json:"formData.values,omitempty"`
	FormDataFile  []FormFileData    `json:"formData.files,omitempty"`
	Explain       []string          `json:"explain"`
}

// 文件类型的数据
type FormFileData struct {
	FormDataField string `json:"formField"`
	FileName      string `json:"fileName"`
	MiniType      string `json:"miniType"`
	Size          int64  `json:"size"`
}

// 处理 echo
func handleEcho(c *gin.Context) {

	// 响应数据
	echoData := EchoData{
		Url:    fmt.Sprintf("%s%s", c.Request.Host, c.Request.URL.String()),
		Method: c.Request.Method,
		Body:   nil,
		Explain: []string{
			"以下是上面的参数说明: ",
			"url: 表示请求地址",
			"method: 表示请求方法",
			"header: 表示请求头",
			"body: 表示请求体数据",
			"query: 表示地址 ? 之后的查询参数",
			"formData.values: 表示 formData 里面的普通数据",
			"formData.files: 表示 formData 里面上传的文件数据",
		},
	}

	// 请求头
	header := make(map[string]string)
	for k, v := range c.Request.Header {
		header[k] = v[0]
	}
	echoData.Header = header

	// fmt.Printf("PostForm: %#v\n", c.Request.PostForm)

	// query 数据
	rawQuery := c.Request.URL.RawQuery
	if len(rawQuery) > 0 {
		query := make(map[string]string)
		for _, q := range strings.Split(rawQuery, "&") {
			items := strings.Split(q, "=")
			query[items[0]] = items[1]
		}
		echoData.Query = query
	}

	// formData 数据
	files, err := c.MultipartForm()
	if err == nil {

		// formData 里的键值对数据
		formDataValue := make(map[string]string)
		for key, value := range files.Value {
			formDataValue[key] = strings.Join(value, ",")
		}
		echoData.FormDataValue = formDataValue

		// formData 里的文件数据
		formDataFile := []FormFileData{}
		for key, value := range files.File {
			// 循环获取文件
			for _, file := range value {
				mimiType := file.Header["Content-Type"][0]
				fd := FormFileData{
					FormDataField: key,           // formData 里面的键
					FileName:      file.Filename, // 文件名
					MiniType:      mimiType,      // 文件的 mimiType
					Size:          file.Size,     // 文件大小
				}
				formDataFile = append(formDataFile, fd)
			}
		}
		echoData.FormDataFile = formDataFile
	}

	// body 数据
	bytes, err := c.GetRawData()
	if err == nil {
		if len(bytes) > 0 {
			body := new(string)
			// 如果 body 数据太长则截取
			if len(bytes) < BODY_DATA_LENGTH {
				*body = string(bytes)
			} else {
				*body = string(bytes[:BODY_DATA_LENGTH]) + "..."
			}
			echoData.Body = body
		}
	}

	c.JSON(http.StatusOK, echoData)
}
