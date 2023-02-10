package userservice

import (
	. "goBlog/db"
	. "goBlog/model"
)

func Signup(user *User) *User {

	return user
}

func GetUserList() (list []User) {
	Db.Find(&list)
	return list
}

func GetUserById(id uint64) (user User) {
	Db.Where("id = ?", id).First(&user)
	return user
}
