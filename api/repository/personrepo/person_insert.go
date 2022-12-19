package personrepo

import (
	"demo-fiber-mysql-gorm/database"
	"demo-fiber-mysql-gorm/model/entity"
)

func (repo *personRepo) InsertPerson_GORM(newPerson *entity.Person) error {

	result := database.Connection.Create(newPerson)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
