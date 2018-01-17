package models

type User struct {
	Id        int
	Email     string
	Telephone string
}

func (user *User) Create() error {
	return DB.Create(user).Error
}

func (user *User) Delete() error {
	return DB.Create(user).Error
}

func UserList(skip int,limit int) ([]*User, error) {
	var users []*User
	err := DB.Limit(limit).Offset(skip).Find(&users).Error
	return users, err
}
