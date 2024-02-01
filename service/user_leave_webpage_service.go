package service

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"singo/model"
	"time"
)

type UserLeaveService struct {
	SerialUUID string
}

func (receiver *UserLeaveService) Operate() {
	if receiver.SerialUUID == "" {
		log.Println("uuid is NULL")
		return
	}
	hash := md5.New()

	// 将字符串转换为字节数组，并计算哈希值
	hash.Write([]byte(receiver.SerialUUID))

	// 将哈希值转换为字符串表示
	serialIDString := hex.EncodeToString(hash.Sum(nil))

	var record model.VisitRecord
	if err := model.DB.Where("serial_uuid = ?", serialIDString).First(&record).Error; err != nil {
		log.Println("handle user leave error: ", err.Error())
		return
	}
	log.Println("[Get Record]: ", record)

	// get record,
	if record.ExplainOpenTime != 0 && record.ExplainCloseTime == 0 {
		record.ExplainCloseTime = time.Now().Unix()
		record.ExplainTime_t = record.ExplainCloseTime - record.ExplainOpenTime
	}

	record.LeaveTime = time.Now().Unix()
	record.VisitTime_t = record.LeaveTime - record.VisitTime

	if err := model.DB.Save(&record).Error; err != nil {
		log.Println("handle user leave error: ", err.Error())
	}
}
