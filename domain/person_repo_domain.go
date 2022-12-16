package domain

import (
	"demo-fiber-mysql-gorm/model/entity"
)

type PersonRepo interface {
	GetPersonWithPersonID_GORM(personId int) (*entity.Person, error)
	InsertPerson_GORM(newPerson *entity.Person) error
	UpdatePerson_GORM(newPerson *entity.Person) error
	DeletePerson_GORM(personID int) error
}
