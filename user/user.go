package user

import (
	"encoding/json"
	"errors"
)

type User struct {
	ID   int64
	Name string
	Tags []string
}

func NewUser(id int64, name string, tags []string) *User {
	return &User{
		ID:   id,
		Name: name,
		Tags: tags,
	}
}


func (user *User) transform() *userInDB {
	tags, _ := json.Marshal(user.Tags)
	return &userInDB{
		ID:   user.ID,
		Name: user.Name,
		Tags: string(tags),
	}
}

func (user *User) IsExist() bool {
	return !db.NewRecord(user.transform())
}

func (user *User) Load() (err error) {
	userDB := loadById(user.ID)
	if userDB == nil {
		err = errors.New("Can't find the user  in database")
		return
	}

	newUser := userDB.transform()
	user.Name = newUser.Name
	user.Tags = newUser.Tags
	return
}

func (user *User) Add() {
	if !user.IsExist() {
		user.transform().Add()
	}
}

func (user *User) hasTags() bool {
	return len(user.Tags) > 0
}



func loadById(id int64) (newUser *userInDB) {
	db.First(newUser, id)
	return
}
