package domain

import "demo-fiber-mysql-gorm/model/entity"

type InnovaRepo interface {
	Get_InnovaOrder_With_InnovaOrderID(innovaOrderID *int, innovaOrder *entity.InnovaOrder) error
	Get_InnovaOrder_With_InnovaOrderID_Slice_WhereIn(conditions *[]int, innovaOrder *[]entity.InnovaOrder) error
	Get_InnovaOrder_With_InnovaOrderID_Slice(SystemBrandID *int, innovaOrder *[]entity.InnovaOrder) error
	Get_InnovaOrder_With_InnovaOrderID_SystemBrandID_AppMemberGroupID(InnovaOrderID, SystemBrandID, AppMemberGroupID *int, innovaOrder *entity.InnovaOrder) error
}
