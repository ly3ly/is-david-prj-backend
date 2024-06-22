package service

import (
	"singo/model"
	"singo/serializer"
)

type UserListPageService struct {
}

func (receiver *UserListPageService) Handler() *serializer.Response {

	var records []model.PageReport
	if err := model.DB.Order("serial_uuid ASC, page_timestamp ASC").Find(&records).Error; err != nil {
		return &serializer.Response{
			Code: 50001,
			Data: err.Error(),
			Msg:  "get records fail!",
		}
	}
	if len(records) == 0 {
		return &serializer.Response{
			Code: 0,
			Msg:  "no record found!",
		}
	}

	type ReportPoint struct {
		Timestamp      int64
		ReportInterval int
		VisitType      int
		Username       string
	}

	PageReportData := map[string][]ReportPoint{}
	ExplainReportData := map[string][]ReportPoint{}

	for i := 0; i < len(records); i++ {
		item := records[i]
		if len(item.SerialUUID) == 0 {
			continue
		}
		if item.PageType == 2 {
			ExplainReportData[item.SerialUUID] = append(ExplainReportData[item.SerialUUID], ReportPoint{
				Timestamp:      item.PageTimestamp,
				ReportInterval: item.ReportInterval,
				VisitType:      item.VisitType,
				Username:       item.UserName,
			})
		}

		PageReportData[item.SerialUUID] = append(PageReportData[item.SerialUUID], ReportPoint{
			Timestamp:      item.PageTimestamp,
			ReportInterval: item.ReportInterval,
			VisitType:      item.VisitType,
			Username:       item.UserName,
		})
	}

	// PageType = 1， 仅页面
	// PageType = 2， 页面+explain
	// explain的时间也计入页面时间

	/*
		页面总时长计算 = 上报点数*interval
		首次进入时间     = 第一个上报点 - interval
		离开时间（近似）  = 最后一个上报点
	*/
	type VisitData struct {
		VisitType      int64  `json:"visit_type"`
		UserName       string `json:"user_name"`
		PageEnterTime  int64  `json:"page_enter_time"`
		PageLeaveTime  int64  `json:"page_leave_time"`
		PageTotalTime  int64  `json:"page_total_time"`
		PageFocusTime  int64  `json:"page_focus_time"`
		PageActiveTime int64  `json:"page_active_time"`
		PageRawPoint   any

		ExplainEnterTime  int64 `json:"explain_enter_time"`
		ExplainLeaveTime  int64 `json:"explain_leave_time"`
		ExplainTotalTime  int64 `json:"explain_total_time"`
		ExplainFocusTime  int64 `json:"explain_focus_time"`
		ExplainActiveTime int64 `json:"explain_active_time"`
		ExplainRawPoint   any
	}

	visitRecord := map[string]VisitData{}
	for key, report := range PageReportData {

		activityService := UserListActivityService{key}
		activityData := activityService.InerHandle()
		timestamps := []int64{}
		for i := 0; i < len(activityData); i++ {
			timestamps = append(timestamps, activityData[i].ActiveTimestamp)
		}
		PageActiveTime := calculateActiveTime(findContinuousTimePoints(timestamps))
		firstReportTime := report[0].Timestamp
		lastReportTime := report[len(report)-1].Timestamp

		var _reportInterval int
		var _visitType int
		var _username string

		if len(report) > 0 {
			_reportInterval = report[0].ReportInterval
			_visitType = report[0].VisitType
			_username = report[0].Username
		}

		var pageEnterTime = firstReportTime - int64(_reportInterval)
		var pageLeaveTime = report[len(report)-1].Timestamp

		if firstReportTime > timestamps[0] {
			pageEnterTime = timestamps[0]
		}

		if lastReportTime < timestamps[len(timestamps)-1] {
			pageLeaveTime = timestamps[len(timestamps)-1]
		}

		visitRecord[key] = VisitData{
			PageEnterTime:  pageEnterTime,
			PageLeaveTime:  pageLeaveTime,
			PageTotalTime:  pageLeaveTime - pageEnterTime,
			PageFocusTime:  int64(len(report) * _reportInterval),
			PageActiveTime: PageActiveTime,
			VisitType:      int64(_visitType),
			UserName:       _username,
			PageRawPoint:   timestamps,
		}
	}

	for key, report := range ExplainReportData {
		visitData := visitRecord[key]

		activityService := UserListActivityService{key}
		PAGE_TYPE := 2 // 2 for explain + page

		activityData := activityService.InerHandleWithType(PAGE_TYPE)
		timestamps := []int64{}
		for i := 0; i < len(activityData); i++ {
			timestamps = append(timestamps, activityData[i].ActiveTimestamp)
		}
		ExplainActiveTime := calculateActiveTime(findContinuousTimePoints(timestamps))

		firstReportTime := report[0].Timestamp
		lastReportTime := report[len(report)-1].Timestamp

		var _reportInterval int

		if len(report) > 0 {
			_reportInterval = report[0].ReportInterval
		}

		var enterTime = firstReportTime - int64(_reportInterval)
		var leaveTime = report[len(report)-1].Timestamp

		if firstReportTime > timestamps[0] {
			enterTime = timestamps[0]
		}

		if lastReportTime < timestamps[len(timestamps)-1] {
			leaveTime = timestamps[len(timestamps)-1]
		}

		visitData.ExplainEnterTime = enterTime
		visitData.ExplainLeaveTime = leaveTime
		visitData.ExplainTotalTime = visitData.ExplainLeaveTime - visitData.ExplainEnterTime
		visitData.ExplainActiveTime = ExplainActiveTime
		visitData.ExplainFocusTime = int64(len(report) * _reportInterval)
		visitData.ExplainRawPoint = timestamps
		visitRecord[key] = visitData
	}

	var rntlist []any

	for key, val := range visitRecord {
		rntlist = append(rntlist, struct {
			SerialID string `json:"serial_id"`
			VisitData
		}{
			SerialID:  key,
			VisitData: val,
		})
	}

	return &serializer.Response{
		Code: 0,
		Data: rntlist,
		Msg:  "get record success!",
	}
}

/*
页面活跃时间，根据activity
非活跃判断： x s内未触发activity
*/
func findContinuousTimePoints(timestamps []int64) [][]int64 {

	const INACTIVE_INTERVAL = 4

	var result [][]int64
	if len(timestamps) == 0 {
		return result
	}

	start := 0
	for i := 1; i < len(timestamps); i++ {
		if timestamps[i]-timestamps[i-1] > INACTIVE_INTERVAL {
			result = append(result, timestamps[start:i])
			start = i
		}
	}
	result = append(result, timestamps[start:])
	return result
}

func calculateActiveTime(times [][]int64) int64 {

	var activeTime int64
	if len(times) == 0 {
		return activeTime
	}

	for i := 0; i < len(times); i++ {
		timeSlice := times[i]
		//if len(timeSlice) < 2 {
		//	activeTime += 2
		//	continue
		//}
		activeTime += timeSlice[len(timeSlice)-1] - timeSlice[0]
	}

	return activeTime
}
