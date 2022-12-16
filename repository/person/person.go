package personrepo

import (
	"demo-fiber-mysql-gorm/domain"

	"gorm.io/gorm"
)

type personRepo struct {
	gorm_db *gorm.DB
}

func NewPersonRepo(gorm_db *gorm.DB) domain.PersonRepo {
	return &personRepo{
		gorm_db: gorm_db,
	}
}
