package user

import (
	"encoding/json"
	"errors"
	"log"
)

type User struct {
	ID   int64
	Name string
	Tags []string
}

type userInDB struct {
	ID   int64  `gorm:"column:id; type:BIGINT(20); PRIMARY_KEY"`
	Name string `gorm:"column:name; type:varchar(100)"`
	Tags string `gorm:"column:tags; type:varchar(200)"`
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

	user.Name = userDB.Name
	if err = json.Unmarshal([]byte(userDB.Tags),
		&user.Tags); err != nil {
		log.Println(err)
	}
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

func (_ *userInDB) TableName() string {
	return "user"
}

func (user *userInDB) transform() *User {
	var tags []string
	_ = json.Unmarshal([]byte(user.Tags), &tags)
	return &User{
		ID:   user.ID,
		Name: user.Name,
		Tags: tags,
	}
}

func (user *userInDB) Add() {
	if db.NewRecord(user) {
		db.Create(user)
	}
}

func loadById(id int64) (newUser *userInDB) {
	db.First(newUser, id)
	return
}
