package innovarepo

import (
	"demo-fiber-mysql-gorm/database"
	"demo-fiber-mysql-gorm/model/entity"
)

// 1 condition with InnovaOrderID
func (innovaRepo *innovaRepo) Get_InnovaOrder_With_InnovaOrderID(innovaOrderID *int, innovaOrder *entity.InnovaOrder) error {

	// get ด้วย primary key
	// err := database.Connection.First(&innovaOrder, "InnovaOrderID = ? AND SystemBrandID = ?", innovaOrderID, SystemBrandID).Error
	// err := database.Connection.First(&innovaOrder, "InnovaOrderID = ?", innovaOrderID).Error
	err := database.Connection.First(&innovaOrder, innovaOrderID).Error // primary key
	if err != nil {
		return err
	}

	return nil
}

// 1 condition with InnovaOrderID Return slice
func (innovaRepo *innovaRepo) Get_InnovaOrder_With_InnovaOrderID_Slice_WhereIn(conditions *[]int, innovaOrder *[]entity.InnovaOrder) error {

	// get ด้วย primary key
	// WHERE in (1,2,3)
	// err := database.Connection.Where(conditions).Find(&innovaOrder).Error // primary key
	err := database.Connection.Find(&innovaOrder, conditions).Error // primary key
	if err != nil {
		return err
	}

	return nil
}

// 1 condition with InnovaOrderID Return slice
func (innovaRepo *innovaRepo) Get_InnovaOrder_With_InnovaOrderID_Slice(SystemBrandID *int, innovaOrder *[]entity.InnovaOrder) error {

	err := database.Connection.Where("SystemBrandID = ?", SystemBrandID).Find(&innovaOrder).Error
	// err := database.Connection.Find(&innovaOrder, SystemBrandID).Error
	if err != nil {
		return err
	}

	return nil
}

// 3 condition with InnovaOrderID, SystemBrandID, AppMemberGroupID
func (innovaRepo *innovaRepo) Get_InnovaOrder_With_InnovaOrderID_SystemBrandID_AppMemberGroupID(InnovaOrderID, SystemBrandID, AppMemberGroupID *int, innovaOrder *entity.InnovaOrder) error {

	err := database.Connection.First(&innovaOrder, InnovaOrderID, SystemBrandID, AppMemberGroupID).Error
	if err != nil {
		return err
	}

	return nil
}
