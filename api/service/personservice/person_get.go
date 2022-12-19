package personservice

import (
	"demo-fiber-mysql-gorm/model/entity"
	"demo-fiber-mysql-gorm/rediscache"
	"encoding/json"
	"fmt"
)

func (service *personService) GetPersonWithPersonID_GORM(personId int) (*entity.Person, error) {

	var person *entity.Person

	// * :: ====================================================
	// * :: Get Redis cache
	personInCache, err := rediscache.Get(fmt.Sprintf("%v", personId))
	if err == nil {
		personMarshal, _ := json.Marshal(personInCache)
		json.Unmarshal([]byte(personMarshal), &person)
		return person, nil
	}
	// * :: ====================================================

	person, err = service.PersonRepo.GetPersonWithPersonID_GORM(personId)
	if err != nil {
		return nil, fmt.Errorf("gorm personID not found")
	}

	return person, nil

}
