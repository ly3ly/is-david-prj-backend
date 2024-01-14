package service

import (
	"singo/model"
	"singo/serializer"
	"time"
)

type UserRecommendServiceV1 struct {
	UserID       int64  `json:"user_id"`
	UserName     string `json:"user_name"`
	SerialUUID	 string `json:"serial_uuid"`
	InputCheck   int64  `json:"input_click"`
	ProcessCheck int64  `json:"process_click"`
	OutputCheck  int64  `json:"output_click"`
}

func (receiver UserRecommendServiceV1) Operate() *serializer.Response {
	var record model.VisitRecord 
	if err := model.DB.Where("serial_uuid = ?", receiver.SerialUUID).First(&record).Error; err != nil {
		return &serializer.Response{
			Code: 50001,
			Data: err.Error(),
			Msg:  "find record fail!",
		}
	}

	record.InputCheck = receiver.InputCheck
	record.ProcessCheck = receiver.ProcessCheck
	record.OutputCheck = receiver.OutputCheck
	record.OperateTime = time.Now().Unix()

	if err := model.DB.Save(&record).Error; err != nil {
		return &serializer.Response{
			Code: 50002,
			Data: err.Error(),
			Msg:  "insert record fail!",
		}
	}

	return &serializer.Response{
		Code: 0,
		Data: record,
		Msg:  "insert record success!",
	}

}
