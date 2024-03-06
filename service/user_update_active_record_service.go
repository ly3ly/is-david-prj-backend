package service

import (
	"crypto/md5"
	"encoding/hex"
	"singo/model"
	"singo/serializer"
)

type UserActiveService struct {
	Type       int    `json:"type"` // 1 page  2 both
	SerialUUID string `json:"serial_uuid" binding:"required"`
	Time       int64  `json:"time"`
}

func (receiver *UserActiveService) Handler() *serializer.Response {
	// 创建一个 MD5 哈希对象
	hash := md5.New()

	// 将字符串转换为字节数组，并计算哈希值
	hash.Write([]byte(receiver.SerialUUID))

	// 将哈希值转换为字符串表示
	serialIDString := hex.EncodeToString(hash.Sum(nil))

	//search SerialUUID from db
	var record model.VisitRecord
	if err := model.DB.Where("serial_uuid = ?", serialIDString).First(&record).Error; err != nil {
		return &serializer.Response{
			Code: 60001,
			Data: err.Error(),
			Msg:  "find record fail!",
		}
	}

	if receiver.Type == 1 {
		record.PageInActiveTime_t += receiver.Time
	}
	if receiver.Type == 2 {
		record.PageInActiveTime_t += receiver.Time
		record.ExplainInActiveTime_t += receiver.Time
	}

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
