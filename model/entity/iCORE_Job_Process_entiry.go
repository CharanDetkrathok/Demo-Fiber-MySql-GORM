package entity

import "time"

type ICOREJobProcess struct {
	Idx           int64     `gorm:"column:idx;primary_key;AUTO_INCREMENT"`
	SystemBrandID int64     `gorm:"column:SystemBrandID"`
	AccountID     int64     `gorm:"column:AccountID;default:0"`
	CardID        int64     `gorm:"column:CardID;default:0"`
	JobName       string    `gorm:"column:JobName"`
	Action        string    `gorm:"column:Action"`
	Value         string    `gorm:"column:Value"`
	Status        string    `gorm:"column:Status;default:Wait"`
	DateCreate    time.Time `gorm:"column:DateCreate"`
	DateModify    time.Time `gorm:"column:DateModify"`
	Comment       string    `gorm:"column:Comment"`
}

func (m *ICOREJobProcess) TableName() string {
	return "iCORE_Job_Process"
}
