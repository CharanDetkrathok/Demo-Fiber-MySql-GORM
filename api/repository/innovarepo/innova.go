package innovarepo

import "demo-fiber-mysql-gorm/api/domain"

type innovaRepo struct{}

func NewInnovaRepo() domain.InnovaRepo {
	return &innovaRepo{}
}
