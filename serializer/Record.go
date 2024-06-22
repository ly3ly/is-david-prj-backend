package serializer

type VisitRecord struct {
	SerialUUID           string
	UserId               int64
	UserName             string
	VisitTime            int64
	LeaveTime            int64
	VisitTime_t          int64
	PageActiveTime_t     int64
	VisitType            int64
	ExplainRecords       []ExplainRecord
	ExplainSumTime       int64
	ExplainSumActiveTime int64
}

type ExplainRecord struct {
	ExplainOpenTime     int64
	ExplainCloseTime    int64
	ExplainTime_t       int64
	ExplainActiveTime_t int64
}
