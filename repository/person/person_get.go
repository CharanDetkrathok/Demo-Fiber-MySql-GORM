package personrepo

import (
	"demo-fiber-mysql-gorm/database"
	"demo-fiber-mysql-gorm/model/entity"
	"fmt"
)

func (repo *personRepo) GetPersonWithPersonID_GORM(personId int) (*entity.Person, error) {
	var person entity.Person
	err := database.Connection.First(&person, personId).Error
	if err != nil {
		fmt.Println("GORM result error :: ", err)
		return &person, err
	}

	return &person, nil
}
