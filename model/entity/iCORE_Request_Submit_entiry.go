package entity

import "time"

type ICORERequestSubmit struct {
	RequestNumber  string    `gorm:"column:RequestNumber;NOT NULL"`
	SystemBrandID  int       `gorm:"column:SystemBrandID;NOT NULL"`
	PaymentRef     string    `gorm:"column:PaymentRef;NOT NULL"`
	PaymentType    string    `gorm:"column:PaymentType"`
	RequestCommand string    `gorm:"column:RequestCommand;NOT NULL"`
	ResultRef      string    `gorm:"column:ResultRef;NOT NULL"`
	DateCreate     time.Time `gorm:"column:DateCreate"`
	RequestID      int64     `gorm:"column:RequestID;default:0"`
	AccountID      int64     `gorm:"column:AccountID;default:0"`
	AccountNumber  string    `gorm:"column:AccountNumber"`
	CardID         int64     `gorm:"column:CardID;default:0"`
	CardNumber     string    `gorm:"column:CardNumber"`
}

func (m *ICORERequestSubmit) TableName() string {
	return "iCORE_Request_Submit"
}
