package service

import (
	"singo/model"
	"singo/serializer"
	"time"
)

type UserRecommendService struct {
	UserID       int64  `json:"user_id"`
	UserName     string `json:"user_name"`
	InputCheck   int64  `json:"input_click"`
	ProcessCheck int64  `json:"process_click"`
	OutputCheck  int64  `json:"output_click"`
}

func (receiver UserRecommendService) Operate() *serializer.Response {
	record := &model.Recommend{
		UserId:       receiver.UserID,
		UserName:     receiver.UserName,
		InputCheck:   receiver.InputCheck,
		ProcessCheck: receiver.ProcessCheck,
		OutputCheck:  receiver.OutputCheck,
		OperateTime:  time.Now().Unix(),
	}
	if err := model.DB.Save(&record).Error; err != nil {
		return &serializer.Response{
			Code: 50001,
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
