package service

import (
	"singo/model"
	"singo/serializer"
	"singo/util"
	"singo/util/e"

	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginServiceV1 struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
}

// Login 用户登录函数
func (service *UserLoginServiceV1) Login(c *gin.Context) serializer.Response {
	var user model.User
	user.UserName = service.UserName

	token, err := util.GenerateToken(service.UserName, service.UserName, 1)
	if err != nil {
		code := e.ERROR_AUTH_TOKEN
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}

	return serializer.BuildUserResponse(user, token)
}
