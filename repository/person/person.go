package personrepo

import (
	"demo-fiber-mysql-gorm/domain"
)

type personRepo struct{}

func NewPersonRepo() domain.PersonRepo {
	return &personRepo{}
}
