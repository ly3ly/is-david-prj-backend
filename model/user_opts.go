package model

type Recommend struct {
	UserId       int64 `gorm:"primaryKey;unique"`
	UserName     string
	InputCheck   int64
	ProcessCheck int64
	OutputCheck  int64
	OperateTime  int64
}
