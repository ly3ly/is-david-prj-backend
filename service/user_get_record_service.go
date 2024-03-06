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
		if item.ExplainOpenTime != 0 && item.ExplainTime_t == 0 {
			continue
		}
		if item.VisitTime_t != 0 && item.LeaveTime == 0 {
			continue
		}
		if item.PageInActiveTime_t > item.VisitTime_t {
			continue
		}
		rntRecords = append(rntRecords, item)
	}

	//得到active的time
	var rnt []model.VisitRecord
	for i := 0; i < len(rntRecords); i++ {
		tmp := rntRecords[i]
		tmp.PageInActiveTime_t = tmp.VisitTime_t - tmp.PageInActiveTime_t
		tmp.ExplainInActiveTime_t = tmp.ExplainTime_t - tmp.ExplainInActiveTime_t
		rnt = append(rnt, tmp)
	}

	return &serializer.Response{
		Code: 0,
		Data: rnt,
		Msg:  "get record success!",
	}

}
