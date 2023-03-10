package personservice

import "demo-fiber-mysql-gorm/api/domain"

type personService struct {
	PersonRepo domain.PersonRepo
}

func NewPersonService(personRepo domain.PersonRepo) domain.PersonService {
	return &personService{
		PersonRepo: personRepo,
	}
}
