package personrepo

import (
	"demo-fiber-mysql-gorm/model/entity"
)

func (repo *personRepo) DeletePerson_GORM(personID int) error {

	result := repo.gorm_db.Where("PersonID = ?", personID).Delete(&entity.Person{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
