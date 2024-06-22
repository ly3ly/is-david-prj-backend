package service

import (
	"singo/model"
	"singo/serializer"
)

type UserListActivityService struct {
	SerialUUID string `json:"serial_uuid" binding:"required"`
}

func (receiver *UserListActivityService) InerHandle() []model.ActiveReport {

	var records []model.ActiveReport
	if err := model.DB.Where("serial_uuid = ?", receiver.SerialUUID).Order("active_timestamp ASC").Find(&records).Error; err != nil {
		//todo

		return records
	}
	return records

}

func (receiver *UserListActivityService) InerHandleWithType(pageType int) []model.ActiveReport {

	var records []model.ActiveReport
	if err := model.DB.Where("serial_uuid = ? and page_type = ?", receiver.SerialUUID, pageType).Order("active_timestamp ASC").Find(&records).Error; err != nil {
		//todo

		return records
	}
	return records

}

func (receiver *UserListActivityService) Handler() *serializer.Response {

	var records []model.ActiveReport
	if err := model.DB.Where("serial_uuid = ?", receiver.SerialUUID).Order("active_timestamp ASC").Find(&records).Error; err != nil {
		return &serializer.Response{
			Code: 50001,
			Data: err.Error(),
			Msg:  "get records fail!",
		}
	}
	//todo

	return &serializer.Response{
		Code: 0,
		Data: records,
		Msg:  "update record success!",
	}
}
