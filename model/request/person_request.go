package request

import (
	"encoding/json"
)

type Person struct {
	PersonID  int    `json:"person_id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Address   string `json:"address"`
	City      string `json:"city"`
}

func (p *Person) Marshal() string {
	personIndent, _ := json.MarshalIndent(p, "", "    ")
	return string(personIndent)
}
