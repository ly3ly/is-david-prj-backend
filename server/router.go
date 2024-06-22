package server

import (
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	//r.Use(middleware.Session("SESSION_SECRET"))
	r.Use(middleware.Cors())
	//r.Use(middleware.CurrentUser())

	//store := cookie.NewStore([]byte("golang-gin:3.1.0"))
	//r.Use(sessions.Sessions("mysession", store))

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		//v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		v1.GET("admin/user/records", api.UserGetRecords)

		// 需要登录保护的
		auth := v1.Group("")
		//auth.Use(middleware.AuthRequired())
		auth.Use(middleware.JWT())
		{
			// User Routing
			//auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
			//auth.PUT("user/subjects", api.UserUpdateSubjects)
			//auth.GET("user/subjects", api.UserGetSubjects)

			auth.PUT("user/records", api.UserUpdateRecords)
			auth.POST("user/inactive", api.UserInActiveService)
			auth.POST("user/activity", api.UserActiveReport)
			auth.POST("user/page", api.UserPageReport)
		}

	}
	return r
}
