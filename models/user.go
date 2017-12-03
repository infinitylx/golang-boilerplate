package models

import (
	"../db"
	"gopkg.in/mgo.v2/bson"
)

// User is user
type User struct {
	UserName  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
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

// SaveUser inserts user into collection
func (h User) SaveUser() error {
	s := db.Clone()
	defer s.Close()

	coll := s.DB("devdb").C("Users")

	err := coll.Insert(bson.M{
		"username":  h.UserName,
		"firstname": h.FirstName,
		"lastname":  h.LastName,
		"password":  h.Password})

	return err

}
