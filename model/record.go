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

	VisitType      int64 // 1 for input-process-output, 2 for input, 3 for process, 4 for output, 5 for none, 6 for input & process, 7 for input & output, 8 for process & output.
	ExplainRecords string
}

type ExplainRecord struct {
	ExplainOpenTime       int64
	ExplainCloseTime      int64
	ExplainTime_t         int64
	ExplainInActiveTime_t int64
}

type ActiveReport struct {
	gorm.Model
	SerialUUID      string
	UserId          int64
	UserName        string
	ActiveTimestamp int64
	ActiveType      string
	PageType        int
	VisitType       int // 1 for input-process-output, 2 for input, 3 for process, 4 for output, 5 for none, 6 for input & process, 7 for input & output, 8 for process & output.
	ExplainOpen     bool
}

type PageReport struct {
	gorm.Model
	SerialUUID     string
	UserId         int64
	UserName       string
	PageTimestamp  int64
	ReportInterval int
	PageType       int
	VisitType      int // 1 for input-process-output, 2 for input, 3 for process, 4 for output, 5 for none, 6 for input & process, 7 for input & output, 8 for process & output.
	ExplainOpen    bool
}
