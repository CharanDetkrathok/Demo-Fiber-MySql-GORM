package entity

import "time"

type ICASHAppMemberGroupDetail struct {
	AppMemberGroupDetailID int64     `gorm:"column:AppMemberGroupDetailID;AUTO_INCREMENT;NOT NULL"`
	AppMemberGroupID       int64     `gorm:"column:AppMemberGroupID;NOT NULL"`
	AccountID              int64     `gorm:"column:AccountID;NOT NULL"`
	AccountNumber          string    `gorm:"column:AccountNumber"`
	DateCreate             time.Time `gorm:"column:DateCreate;default:CURRENT_TIMESTAMP;NOT NULL"`
	DateModify             time.Time `gorm:"column:DateModify;default:CURRENT_TIMESTAMP;NOT NULL"`
	AdminIDModify          int64     `gorm:"column:AdminIDModify;default:0"`
}

func (m *ICASHAppMemberGroupDetail) TableName() string {
	return "iCASH_AppMemberGroupDetail"
}
