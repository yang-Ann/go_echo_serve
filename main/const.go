package main

var DefaultLanguage = "zh-CN"
var HeaderLanguageKey = "Language"

// 参数说明
var ExplainMap = map[string][]string{
	"zh-CN": {
		"以下是上面的参数说明：",
		"url：表示请求地址",
		"method：表示请求方法",
		"header：表示请求头",
		"body：表示请求体数据",
		"query：表示地址 ? 之后的查询参数",
		"formData.values：表示 formData 里面的普通数据",
		"formData.files：表示 formData 里面上传的文件数据",
	},
	"en-US": {
		"The following is a description of the above parameters:",
		"url: indicates the request address",
		"method: Indicates the request method",
		"header: Indicates the request header",
		"body: indicates the request body data",
		"query: indicates the address? Subsequent query parameters",
		"Formdata. values: common data in formData",
		"Formdata. files: Indicates the file data uploaded in formData",
	},
}