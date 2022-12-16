package personhandle

import (
	"demo-fiber-mysql-gorm/domain"
)

type personHandle struct {
	PersonService domain.PersonService
}

func NewPersonHandle(PersonService domain.PersonService) domain.PersonHandle {
	return &personHandle{PersonService: PersonService}
}
