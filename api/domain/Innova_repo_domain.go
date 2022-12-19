package domain

import "demo-fiber-mysql-gorm/model/entity"

type InnovaRepo interface {
	Get_InnovaOrder_With_InnovaOrderID(innovaOrderID *int, innovaOrder *entity.InnovaOrder) error
}
