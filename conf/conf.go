package conf

import (
	"os"
	"singo/model"
	"singo/util"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	//if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
	//	util.Log().Panic("翻译文件加载失败", err)
	//}

	// 连接数据库
	// model.Database(os.Getenv("MYSQL_DSN"))
	// model.Database("root:flora9966@tcp(localhost:3306)/recruitmentdb?charset=utf8mb4&parseTime=True&loc=Local")

	// model.Database("root:flora9966@tcp(127.0.0.1:3306)/recruitmentdb?charset=utf8mb4&parseTime=True&loc=Local")
	model.Database("root:5KW2sWIxkQq7X9c@tcp(127.0.0.1:3306)/recruitmentdb?charset=utf8mb4&parseTime=True&loc=Local")
	//cache.Redis()
}
