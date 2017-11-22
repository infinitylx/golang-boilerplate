package models

import (
	"github.com/faerulsalamun/golang-boilerplate/db"
)

type User struct {
	Name string `json:"name"`
}

func (h User) GetAll() ([] User, error) {

	s := db.Clone()
	defer s.CloseSession()

	var result [] User
	err := s.DB("devdb").c("Users").Find(nil).All(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
