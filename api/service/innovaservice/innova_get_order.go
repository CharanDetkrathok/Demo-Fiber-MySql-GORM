package innovaservice

import (
	"demo-fiber-mysql-gorm/model/entity"
	"fmt"
)

// 1 condition with InnovaOrderID
func (innovaService *innovaService) Get_InnovaOrder_With_InnovaOrderID(InnovaOrderID *int) (*entity.InnovaOrder, error) {

	innovaOrder := entity.InnovaOrder{}
	err := innovaService.InnovaRepo.Get_InnovaOrder_With_InnovaOrderID(InnovaOrderID, &innovaOrder)
	if err != nil {
		fmt.Println("Get Innova Order failed ERROR :: ", err)
		return nil, err
	}

	fmt.Println("Get Innova Order SUCCESS :: ", innovaOrder)
	return &innovaOrder, nil
}

// 1 condition with InnovaOrderID Slice
func (innovaService *innovaService) Get_InnovaOrder_With_InnovaOrderID_Slice_WhereIn(InnovaOrderID1, InnovaOrderID2, InnovaOrderID3 *int) (*[]entity.InnovaOrder, error) {

	innovaOrder := []entity.InnovaOrder{}
	conditions := []int{*InnovaOrderID1, *InnovaOrderID2, *InnovaOrderID3}
	err := innovaService.InnovaRepo.Get_InnovaOrder_With_InnovaOrderID_Slice_WhereIn(&conditions, &innovaOrder)
	if err != nil {
		fmt.Println("Get Innova Order failed ERROR :: ", err)
		return nil, err
	}

	fmt.Println("Get Innova Order SUCCESS :: ", innovaOrder)
	return &innovaOrder, nil
}

// 1 condition with InnovaOrderID Slice
func (innovaService *innovaService) Get_InnovaOrder_With_InnovaOrderID_Slice(SystemBrandID *int) (*[]entity.InnovaOrder, error) {

	innovaOrder := []entity.InnovaOrder{}
	err := innovaService.InnovaRepo.Get_InnovaOrder_With_InnovaOrderID_Slice(SystemBrandID, &innovaOrder)
	if err != nil {
		fmt.Println("Get Innova Order failed ERROR :: ", err)
		return nil, err
	}

	fmt.Println("Get Innova Order SUCCESS :: ", innovaOrder)
	return &innovaOrder, nil
}

// 3 condition with InnovaOrderID, SystemBrandID, AppMemberGroupID
func (innovaService *innovaService) Get_InnovaOrder_With_InnovaOrderID_SystemBrandID_AppMemberGroupID(InnovaOrderID, SystemBrandID, AppMemberGroupID *int) (*entity.InnovaOrder, error) {

	innovaOrder := entity.InnovaOrder{}
	err := innovaService.InnovaRepo.Get_InnovaOrder_With_InnovaOrderID_SystemBrandID_AppMemberGroupID(InnovaOrderID, SystemBrandID, AppMemberGroupID, &innovaOrder)
	if err != nil {
		fmt.Println("Get Innova Order failed ERROR :: ", err)
		return nil, err
	}

	fmt.Println("Get Innova Order SUCCESS :: ", innovaOrder)
	return &innovaOrder, nil
}
