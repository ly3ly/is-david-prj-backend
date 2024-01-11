package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func UserUpdateSubjects(c *gin.Context) {
	var service service.UserRecommendService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Operate()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserGetSubjects(c *gin.Context) {
	var service service.UserGetSubjectsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Operate()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
