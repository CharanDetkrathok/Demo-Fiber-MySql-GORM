package domain

import (
	"demo-fiber-mysql-gorm/model/entity"
	"demo-fiber-mysql-gorm/model/request"
)

type PersonService interface {
	GetPersonWithPersonID_GORM(personId int) (*entity.Person, error)
	InsertPerson_GORM(newPerson *request.Person) error
	UpdatePerson_GORM(newPerson *request.Person) error
	DeletePerson_GORM(personID int) error
}
