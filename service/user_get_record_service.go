package service

import (
	"encoding/json"
	"fmt"
	"singo/model"
	"singo/serializer"
)

type UserGetRecordsService struct {
}

/*
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
*/

func (receiver UserGetRecordsService) OperateV1() *serializer.Response {
	var records []model.VisitRecord
	if err := model.DB.Find(&records).Error; err != nil {
		return &serializer.Response{
			Code: 50001,
			Data: err.Error(),
			Msg:  "get records fail!",
		}
	}

	var rntRecords []serializer.VisitRecord
	for i := 0; i < len(records); i++ {
		item := records[i]
		if len(item.SerialUUID) == 0 {
			continue
		}
		if item.VisitTime_t != 0 && item.LeaveTime == 0 {
			continue
		}
		if item.PageInActiveTime_t > item.VisitTime_t {
			continue
		}

		var explainRecords []model.ExplainRecord
		err := json.Unmarshal([]byte(item.ExplainRecords), &explainRecords)
		if err != nil {
			fmt.Println("error in get explain records: ", item.ExplainRecords)
		}
		var rntExplainRecords []serializer.ExplainRecord
		var explainSumActiveTime int64 = 0
		var explainSumTime int64 = 0
		for j := 0; j < len(explainRecords); j++ {
			explainRecord := explainRecords[j]
			if explainRecord.ExplainTime_t > item.VisitTime_t {
				continue
			}
			if explainRecord.ExplainOpenTime != 0 && explainRecord.ExplainTime_t == 0 {
				continue
			}
			explainActiveTime := explainRecord.ExplainTime_t - explainRecord.ExplainInActiveTime_t
			rntExplainRecords = append(rntExplainRecords, serializer.ExplainRecord{
				ExplainOpenTime:     explainRecord.ExplainOpenTime,
				ExplainCloseTime:    explainRecord.ExplainCloseTime,
				ExplainTime_t:       explainRecord.ExplainTime_t,
				ExplainActiveTime_t: explainActiveTime,
			})
			explainSumActiveTime += explainActiveTime
			explainSumTime += explainRecord.ExplainTime_t
		}
		pageActiveTime := item.VisitTime_t - item.PageInActiveTime_t
		if explainSumActiveTime > pageActiveTime {
			continue
		}
		if explainSumTime > item.VisitTime_t {
			continue
		}
		rntRecords = append(rntRecords, serializer.VisitRecord{
			SerialUUID:           item.SerialUUID,
			UserId:               item.UserId,
			UserName:             item.UserName,
			VisitTime:            item.VisitTime,
			LeaveTime:            item.LeaveTime,
			VisitTime_t:          item.VisitTime_t,
			PageActiveTime_t:     pageActiveTime,
			VisitType:            item.VisitType,
			ExplainRecords:       rntExplainRecords,
			ExplainSumTime:       explainSumTime,
			ExplainSumActiveTime: explainSumActiveTime,
		})
	}

	return &serializer.Response{
		Code: 0,
		Data: rntRecords,
		Msg:  "get record success!",
	}
}
