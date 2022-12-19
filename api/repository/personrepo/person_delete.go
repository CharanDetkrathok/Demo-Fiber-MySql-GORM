package personrepo

import (
	"demo-fiber-mysql-gorm/database"
	"demo-fiber-mysql-gorm/model/entity"
)

func (repo *personRepo) DeletePerson_GORM(personID int) error {

	result := database.Connection.Where("PersonID = ?", personID).Delete(&entity.Person{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
