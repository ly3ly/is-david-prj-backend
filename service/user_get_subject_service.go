package service

import (
	"singo/model"
	"singo/serializer"
)

type UserGetSubjectsService struct {
}

func (receiver UserGetSubjectsService) Operate() *serializer.Response {
	var records []model.Recommend

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
