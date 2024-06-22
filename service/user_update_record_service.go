package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"singo/model"
	"singo/serializer"
	"time"
)

type UserUpdateRecordService struct {
	UserID     int64  `json:"user_id" binding:"required"`
	UserName   string `json:"user_name"`
	VisitType  int64  `json:"visit_type"`                     // 1 for input-process-output, 2 for input, 3 for process, 4 for output, 5 for none
	SerialUUID string `json:"serial_uuid" binding:"required"` //distribute while TimeType = 1
	TimeType   int64  `json:"time_type" binding:"required"`   // 1: visit_time, 2:leave_time, 3: explain_open_time, 4: explain_close_time
	// Time     	 int64   `json:"time"`
}

/*
func (receiver UserUpdateRecordService) Operate() *serializer.Response {

	//serialID := md5.Sum([]byte(receiver.SerialUUID)) //token as serial id
	//serialIDString := fmt.Sprintf("%v", serialID)
	//serialIDString, _ := bcrypt.GenerateFromPassword([]byte(receiver.SerialUUID), 5)

	// 创建一个 MD5 哈希对象
	hash := md5.New()

	// 将字符串转换为字节数组，并计算哈希值
	hash.Write([]byte(receiver.SerialUUID))

	// 将哈希值转换为字符串表示
	serialIDString := hex.EncodeToString(hash.Sum(nil))

	fmt.Println("Receive time_type: ", receiver.TimeType)
	if receiver.TimeType == 1 {
		record := &model.VisitRecord{
			UserId:   receiver.UserID,
			UserName: receiver.UserName,
			//SerialUUID: uuid.NewString(),
			SerialUUID: serialIDString,
			VisitTime:  time.Now().Unix(),
			VisitType:  receiver.VisitType,
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

	//if TimeType != 1
	//search SerialUUID from db
	var record model.VisitRecord
	if err := model.DB.Where("serial_uuid = ?", serialIDString).First(&record).Error; err != nil {
		return &serializer.Response{
			Code: 50002,
			Data: err.Error(),
			Msg:  "find record fail!",
		}
	}

	// get record,
	// if TimeType = 4 && ExplainOpenTime missing ==> ignore this request ==> delete this record from db
	//if TimeType = 2 && LeaveTime missing ==> ignore this request ==> delete this record from db
	if receiver.TimeType == 4 && record.ExplainOpenTime == 0 ||
		receiver.TimeType == 2 && record.VisitTime == 0 {
		return &serializer.Response{
			Code: 50003,
			Data: record,
			Msg:  "error in locating starting time, will dismiss this record...",
		}
	}

	if receiver.TimeType == 2 {
		record.LeaveTime = time.Now().Unix()
		record.VisitTime_t = record.LeaveTime - record.VisitTime
	}

	if receiver.TimeType == 3 {
		if record.ExplainOpenTime == 0 {
			record.ExplainOpenTime = time.Now().Unix()
		}
	}

	if receiver.TimeType == 4 {
		record.ExplainCloseTime = time.Now().Unix()
		record.ExplainTime_t = record.ExplainTime_t + record.ExplainCloseTime - record.ExplainOpenTime
	}

	if err := model.DB.Save(&record).Error; err != nil {
		return &serializer.Response{
			Code: 50004,
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
*/

func (receiver UserUpdateRecordService) OperateV1() *serializer.Response {
	// 创建一个 MD5 哈希对象
	hash := md5.New()

	// 将字符串转换为字节数组，并计算哈希值
	hash.Write([]byte(receiver.SerialUUID))

	// 将哈希值转换为字符串表示
	serialIDString := hex.EncodeToString(hash.Sum(nil))

	fmt.Println("Receive time_type: ", receiver.TimeType)
	if receiver.TimeType == 1 {
		explainRecords, err := json.Marshal([]model.ExplainRecord{})
		if err != nil {
			fmt.Println("init empty explain error: ", err.Error())
			explainRecords = []byte("[]")
		}
		record := &model.VisitRecord{
			UserId:   receiver.UserID,
			UserName: receiver.UserName,
			//SerialUUID: uuid.NewString(),
			SerialUUID:     serialIDString,
			VisitTime:      time.Now().Unix(),
			VisitType:      receiver.VisitType,
			ExplainRecords: string(explainRecords),
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
	//if TimeType != 1
	//search SerialUUID from db
	var record model.VisitRecord
	if err := model.DB.Where("serial_uuid = ?", serialIDString).First(&record).Error; err != nil {
		return &serializer.Response{
			Code: 50002,
			Data: err.Error(),
			Msg:  "find record fail!",
		}
	}

	if receiver.TimeType == 2 {
		if record.VisitTime == 0 {
			return &serializer.Response{
				Code: 50004,
				Data: record,
				Msg:  "error in locating starting time, will dismiss this record...",
			}
		}
		record.LeaveTime = time.Now().Unix()
		record.VisitTime_t = record.LeaveTime - record.VisitTime
	}

	var explainRecords []model.ExplainRecord
	fmt.Println("record.ExplainRecords: ", record.ExplainRecords)
	err := json.Unmarshal([]byte(record.ExplainRecords), &explainRecords)
	if err != nil {
		return &serializer.Response{
			Code: 50003,
			Data: err.Error(),
			Msg:  "extract explain records fail!",
		}
	}
	if receiver.TimeType == 3 {
		explainRecords = append(explainRecords, model.ExplainRecord{
			ExplainOpenTime:       time.Now().Unix(),
			ExplainCloseTime:      0,
			ExplainTime_t:         0,
			ExplainInActiveTime_t: 0,
		})
	}

	if receiver.TimeType == 4 {
		lastExplainRecord := explainRecords[len(explainRecords)-1]
		if lastExplainRecord.ExplainOpenTime == 0 {
			return &serializer.Response{
				Code: 50004,
				Data: lastExplainRecord,
				Msg:  "error in locating starting time, will dismiss this record...",
			}
		}
		lastExplainRecord.ExplainCloseTime = time.Now().Unix()
		lastExplainRecord.ExplainTime_t = lastExplainRecord.ExplainCloseTime - lastExplainRecord.ExplainOpenTime
		explainRecords[len(explainRecords)-1] = lastExplainRecord
	}

	explainRecordsBytes, _ := json.Marshal(explainRecords)
	record.ExplainRecords = string(explainRecordsBytes)

	if err := model.DB.Save(&record).Error; err != nil {
		return &serializer.Response{
			Code: 50005,
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
