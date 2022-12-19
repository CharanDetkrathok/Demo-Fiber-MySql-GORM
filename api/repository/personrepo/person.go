package personrepo

import "demo-fiber-mysql-gorm/api/domain"

type personRepo struct{}

func NewPersonRepo() domain.PersonRepo {
	return &personRepo{}
}
