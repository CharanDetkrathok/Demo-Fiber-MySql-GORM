package entity

import "time"

type InnovaOrder struct {
	InnovaOrderID      int64     `gorm:"column:InnovaOrderID;AUTO_INCREMENT;NOT NULL"`
	SystemBrandID      int64     `gorm:"column:SystemBrandID"`
	AppMemberGroupID   int64     `gorm:"column:AppMemberGroupID"`
	AccountID          int64     `gorm:"column:AccountID"`
	CardNumber         string    `gorm:"column:CardNumber"`
	RequestID          int64     `gorm:"column:RequestID"`     // RequestID ของ PromotionUse
	RequestNumber      string    `gorm:"column:RequestNumber"` // RequestNumber ของ PromotionUse
	OrderNumber        string    `gorm:"column:OrderNumber"`
	OrderTotalAmount   string    `gorm:"column:OrderTotalAmount"`
	RewardCode         string    `gorm:"column:RewardCode"`
	AssignRewardStatus string    `gorm:"column:AssignRewardStatus"` // Wait, Success ,Fail
	DateAssignReward   time.Time `gorm:"column:DateAssignReward"`   // วันที่เรียก Innova success
	DateOrder          time.Time `gorm:"column:DateOrder"`
	DateCreate         time.Time `gorm:"column:DateCreate"`
	DateModify         time.Time `gorm:"column:DateModify"`
	Reason             string    `gorm:"column:Reason"`
}

func (m *InnovaOrder) TableName() string {
	return "Innova_Order"
}
