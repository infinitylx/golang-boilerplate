package models

import (
	"../db"
)

// User is user
type User struct {
	Name string `json:"name"`
}

// GetAll return list of all users
func (h User) GetAll() ([]User, error) {

	s := db.Clone()
	defer s.Close()

	var result []User
	err := s.DB("devdb").C("Users").Find(nil).All(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
