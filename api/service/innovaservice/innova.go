package innovaservice

import "demo-fiber-mysql-gorm/api/domain"

type innovaService struct {
	InnovaRepo domain.InnovaRepo
}

func NewInnovaService(innovaRepo domain.InnovaRepo) domain.InnovaService {
	return &innovaService{InnovaRepo: innovaRepo}
}
