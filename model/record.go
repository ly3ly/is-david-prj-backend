package model

import (
	"gorm.io/gorm"
)

type VisitRecord struct {
	gorm.Model
	SerialUUID         string `gorm:"unique"`
	UserId             int64
	UserName           string
	VisitTime          int64
	LeaveTime          int64
	VisitTime_t        int64
	PageInActiveTime_t int64

	VisitType             int64 // 1 for input-process-output, 2 for input, 3 for process, 4 for output, 5 for none
	ExplainOpenTime       int64
	ExplainCloseTime      int64
	ExplainTime_t         int64
	ExplainInActiveTime_t int64
}
