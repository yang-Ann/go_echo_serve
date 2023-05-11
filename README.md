## 启动服务

- window
  - 点击 `main.exe`即可(默认端口8100)
  - 命令行启动: `main.exe -port 8200`




## 本地开发

```sh

# 下载依赖
go mod tidy

# 启动服务(默认监听8100)
go run .

# 指定服务的端口
go run . -port=8200

#------------------- 使用 fresh 启动(文件更新自动重启服务)
# 安装 fresh
go install github.com/pilu/fresh

# 启动
fresh
```