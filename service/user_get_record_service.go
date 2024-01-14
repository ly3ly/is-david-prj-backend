package service

import (
	"singo/model"
	"singo/serializer"
)

type UserGetRecordsService struct {
}

func (receiver UserGetRecordsService) Operate() *serializer.Response {
	var records []model.VisitRecord

	if err := model.DB.Find(&records).Error; err != nil {
		return &serializer.Response{
			Code: 50001,
			Data: err.Error(),
			Msg:  "get records fail!",
		}
	}

	return &serializer.Response{
		Code: 0,
		Data: records,
		Msg:  "get record success!",
	}

}
