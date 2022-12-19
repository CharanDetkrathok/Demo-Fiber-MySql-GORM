package innovaservice

import (
	"demo-fiber-mysql-gorm/model/entity"
	"fmt"
)

func (innovaService *innovaService) Get_InnovaOrder_With_InnovaOrderID(innovaOrderID *int) (*entity.InnovaOrder, error) {

	innovaOrder := entity.InnovaOrder{}
	err := innovaService.InnovaRepo.Get_InnovaOrder_With_InnovaOrderID(innovaOrderID, &innovaOrder)
	if err != nil {
		fmt.Println("Get Innova Order failed ERROR :: ", err)
		return nil, err
	}

	fmt.Println("Get Innova Order SUCCESS :: ", innovaOrder)
	return &innovaOrder, nil
}
