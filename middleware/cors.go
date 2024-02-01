package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"regexp"
)

// Cors 跨域配置
//func Cors() gin.HandlerFunc {
//	config := cors.DefaultConfig()
//	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
//	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
//	if gin.Mode() == gin.ReleaseMode {
//		// 生产环境需要配置跨域域名，否则403
//		config.AllowOrigins = []string{"http://www.example.com"}
//	} else {
//		// 测试环境下模糊匹配本地开头的请求
//		config.AllowOriginFunc = func(origin string) bool {
//			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
//				return true
//			}
//			if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
//				return true
//			}
//			return false
//		}
//	}
//	config.AllowCredentials = true
//	return cors.New(config)
//}
//
//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		origin := c.Request.Header.Get("Origin")
//		if origin != "" {
//			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
//			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
//			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
//			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
//			c.Header("Access-Control-Allow-Credentials", "true")
//		}
//		if method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusNoContent)
//		}
//		c.Next()
//	}
//}

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "Authorization"}
	// config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:8081", "http://j9grvkwx.shenzhuo.vip:12131","http://localhost:3000","http://localhost:3001","http://localhost:3002","http://192.168.2.120:8080","http://192.168.2.120:12345","http://192.168.2.120:3001","http://192.168.2.120:3002"}

	// 测试环境下模糊匹配本地开头的请求
	config.AllowOriginFunc = func(origin string) bool {
		if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
			return true
		}
		if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
			return true
		}

		return true
	}

	config.AllowCredentials = true
	return cors.New(config)
}
