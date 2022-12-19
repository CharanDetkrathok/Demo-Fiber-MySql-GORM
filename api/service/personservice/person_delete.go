package personservice

import (
	"fmt"
)

func (service *personService) DeletePerson_GORM(personID int) error {

	_, err := service.PersonRepo.GetPersonWithPersonID_GORM(personID)
	if err != nil {
		fmt.Println("gorm delete failed, not found person id. :: ", err)
		return fmt.Errorf("gorm delete failed, not found person id")
	}

	err = service.PersonRepo.DeletePerson_GORM(personID)
	if err != nil {
		fmt.Println("gorm delete failed. :: ", err)
		return fmt.Errorf("gorm delete failed")
	}

	fmt.Println("SUCCESS delete by GORM")
	return nil

}
