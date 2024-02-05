# Singo

Singo: Simple Single Golang Web Service

go-crud正式改名为Singo!

使用Singo开发Web服务: 用最简单的架构，实现够用的框架，服务海量用户

https://github.com/Gourouting/singo


## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init singo
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```

## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)
