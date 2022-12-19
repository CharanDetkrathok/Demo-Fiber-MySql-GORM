package domain

import "demo-fiber-mysql-gorm/model/entity"

type InnovaService interface {
	Get_InnovaOrder_With_InnovaOrderID(innovaOrderID *int) (*entity.InnovaOrder, error)
}
