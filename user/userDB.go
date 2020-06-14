package user

import "encoding/json"

type userInDB struct {
	ID   int64  `gorm:"column:id; type:BIGINT(20); PRIMARY_KEY"`
	Name string `gorm:"column:name; type:varchar(100)"`
	Tags string `gorm:"column:tags; type:varchar(200)"`
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
