package innovarepo

import (
	"demo-fiber-mysql-gorm/database"
	"demo-fiber-mysql-gorm/model/entity"
)

func (innovaRepo *innovaRepo) Get_InnovaOrder_With_InnovaOrderID(innovaOrderID *int, innovaOrder *entity.InnovaOrder) error {

	err := database.Connection.First(&innovaOrder, innovaOrderID).Error
	if err != nil {
		return err
	}

	return nil
}
