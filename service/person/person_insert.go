package personservice

import (
	"demo-fiber-mysql-gorm/model/entity"
	"demo-fiber-mysql-gorm/model/request"
	"fmt"
)

func (service *personService) InsertPerson_GORM(newPerson *request.Person) error {

	newPersonTemp := entity.Person{
		PersonID:  newPerson.PersonID,
		LastName:  newPerson.LastName,
		FirstName: newPerson.FirstName,
		Address:   newPerson.Address,
		City:      newPerson.City,
	}

	err := service.PersonRepo.InsertPerson_GORM(&newPersonTemp)
	if err != nil {
		fmt.Println("INSERT GORM failed. :: ", err)
		return fmt.Errorf("gorm insert failed")
	}

	fmt.Println("SUCCESS Insert by GORM")
	return nil

}
