package entity

import "time"

type ICASHCardActivate struct {
	CardID                       int64     `gorm:"column:CardID;NOT NULL"`
	SystemBrandID                int       `gorm:"column:SystemBrandID"`
	AccountID                    int64     `gorm:"column:AccountID"`
	CardBillingNumber            string    `gorm:"column:CardBillingNumber"` // Unique Account Number for PaymentGateway
	CardNumber                   string    `gorm:"column:CardNumber;NOT NULL"`
	CardTypeID                   int       `gorm:"column:CardTypeID;NOT NULL"`
	CardStatus                   string    `gorm:"column:CardStatus"` // New/Active/Inactive/Terminated/Block
	DateExpire                   time.Time `gorm:"column:DateExpire"`
	DateCreate                   time.Time `gorm:"column:DateCreate"`
	AdminIDCreate                int64     `gorm:"column:AdminIDCreate"`
	DateModify                   time.Time `gorm:"column:DateModify"`
	AdminIDModify                int64     `gorm:"column:AdminIDModify"`
	RequestID                    int64     `gorm:"column:RequestID"`
	IsAllowDeactivate            int       `gorm:"column:IsAllowDeactivate;default:1"`
	RequestIDCreate              int64     `gorm:"column:RequestIDCreate;default:0"`
	CardRefCode                  string    `gorm:"column:CardRefCode"`
	CardGraphicCode              string    `gorm:"column:CardGraphicCode"`
	IsAllowRequest               int       `gorm:"column:IsAllowRequest;default:1"`
	IsAllowRequestCardPresent    int       `gorm:"column:IsAllowRequestCardPresent;default:1"`
	IsAllowRequestCardNotPresent int       `gorm:"column:IsAllowRequestCardNotPresent;default:1"`
	CardSerialNumber             string    `gorm:"column:CardSerialNumber"`
	CreditTopupProductSettingID  string    `gorm:"column:CreditTopupProductSettingID;default:0"` // CreditTopupProductSettingID for Topup
}

func (m *ICASHCardActivate) TableName() string {
	return "iCASH_Card_Activate"
}
