package personrepo

import (
	"demo-fiber-mysql-gorm/model/entity"
)

func (repo *personRepo) InsertPerson_GORM(newPerson *entity.Person) error {

	result := repo.gorm_db.Create(newPerson)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
