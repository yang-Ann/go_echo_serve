## 启动

```sh

# 下载依赖
go mod tidy

cd main

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

## 使用

```sh
curl localhost:8100 -G -d name=zhangsan -d age=18

echo hello > data.txt
curl -F file=@data.txt -F age=18 localhost:8100
```

## 打包
```sh
cd main

go build

# 启动服务(默认监听8100)
.\main.exe

# 指定服务的端口
.\main.exe -port=8200
```