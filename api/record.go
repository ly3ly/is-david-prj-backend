package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

func UserUpdateRecords(c *gin.Context) {
	var service service.UserUpdateRecordService
	if err := c.ShouldBind(&service); err == nil {
		res := service.OperateV1()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserGetRecords(c *gin.Context) {
	var service service.UserListPageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Handler()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserInActiveService(c *gin.Context) {
	var service service.UserActiveService
	if err := c.ShouldBind(&service); err == nil {
		res := service.HandlerV1()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserPageReport(c *gin.Context) {
	var service service.UserPageReportService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Handler()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserActiveReport(c *gin.Context) {
	var service service.UserActiveReportService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Handler()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
