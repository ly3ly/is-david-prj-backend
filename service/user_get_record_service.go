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

	rntRecords := []model.VisitRecord{}

	for i := 0; i < len(records); i++ {
		item := records[i]
		if len(item.SerialUUID) == 0 {
			continue
		}
		if item.ExplainTime_t > item.VisitTime_t {
			continue
		}
		rntRecords = append(rntRecords, item)
	}

	return &serializer.Response{
		Code: 0,
		Data: rntRecords,
		Msg:  "get record success!",
	}

}
