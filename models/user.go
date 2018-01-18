package models

import "time"

type User struct {
	Id         int
	Email      string
	Salt       string
	Password   string
	IsActived  bool
	IsBlocked  bool
	Createtime time.Time
	Updatetime time.Time
}

func (user *User) Create() error {
	return DB.Create(user).Error
}

func (user *User) Delete() error {
	return DB.Delete(user).Error
}

func (user *User) First() error {
	return DB.First(user).Error
}

func UserList(skip int, limit int) ([]User, error) {
	var users []User
	err := DB.Limit(limit).Offset(skip).Find(&users).Error
	return users, err
}
