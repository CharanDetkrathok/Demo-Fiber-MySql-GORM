package domain

import "demo-fiber-mysql-gorm/model/entity"

type InnovaService interface {
	Get_InnovaOrder_With_InnovaOrderID(innovaOrderID *int) (*entity.InnovaOrder, error)
	Get_InnovaOrder_With_InnovaOrderID_Slice_WhereIn(InnovaOrderID1, InnovaOrderID2, InnovaOrderID3 *int) (*[]entity.InnovaOrder, error)
	Get_InnovaOrder_With_InnovaOrderID_Slice(SystemBrandID *int) (*[]entity.InnovaOrder, error)
	Get_InnovaOrder_With_InnovaOrderID_SystemBrandID_AppMemberGroupID(InnovaOrderID, SystemBrandID, AppMemberGroupID *int) (*entity.InnovaOrder, error)
}
