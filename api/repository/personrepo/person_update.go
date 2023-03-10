package personrepo

import (
	"demo-fiber-mysql-gorm/database"
	"demo-fiber-mysql-gorm/model/entity"
)

func (repo *personRepo) UpdatePerson_GORM(newPerson *entity.Person) error {

	result := database.Connection.Model(&newPerson).Where("PersonID = ?", newPerson.PersonID).Updates(newPerson)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
