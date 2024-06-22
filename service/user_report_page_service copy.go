package service

import (
	"crypto/md5"
	"encoding/hex"
	"singo/model"
	"singo/serializer"
	"time"
)

type UserPageReportService struct {
	UserName    string `json:"user_name"`
	VisitType   int    `json:"visit_type"`
	SerialUUID  string `json:"serial_uuid" binding:"required"`
	Time        int64  `json:"time"`
	Interval    int    `json:"interval"`
	PageType    int    `json:"page_type"`
	ExplainOpen bool   `json:"explain_open"`
}

func (receiver *UserPageReportService) Handler() *serializer.Response {
	// 创建一个 MD5 哈希对象
	hash := md5.New()

	// 将字符串转换为字节数组，并计算哈希值
	hash.Write([]byte(receiver.SerialUUID))

	// 将哈希值转换为字符串表示
	serialIDString := hex.EncodeToString(hash.Sum(nil))

	//search SerialUUID from db
	var record model.PageReport

	record.UserName = receiver.UserName
	record.VisitType = receiver.VisitType
	record.SerialUUID = serialIDString
	record.PageTimestamp = time.Now().Unix()
	record.ReportInterval = receiver.Interval
	record.PageType = receiver.PageType
	record.ExplainOpen = receiver.ExplainOpen

	if err := model.DB.Save(&record).Error; err != nil {
		return &serializer.Response{
			Code: 60002,
			Data: err.Error(),
			Msg:  "update record fail!",
		}
	}

	return &serializer.Response{
		Code: 0,
		Data: receiver,
		Msg:  "update record success!",
	}
}
