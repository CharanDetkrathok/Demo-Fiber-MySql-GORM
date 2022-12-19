package personhandler

import "demo-fiber-mysql-gorm/api/domain"

type personHandler struct {
	PersonService domain.PersonService
}

func NewPersonHandler(PersonService domain.PersonService) domain.PersonHandler {
	return &personHandler{PersonService: PersonService}
}
