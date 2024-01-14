package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func UserUpdateRecords(c *gin.Context) {
	var service service.UserUpdateRecordService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Operate()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserGetRecords(c *gin.Context) {
	var service service.UserGetRecordsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Operate()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

