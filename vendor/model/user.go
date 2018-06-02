package model

import (
	"database"
	"entity"
)

// GetUserByID ...
func GetUserByID(userid int64) (user *entity.User, has bool, er error) {
	user = &entity.User{ID: userid}
	has, er = database.Orm.Table("user").Get(user)
	return
}

// GetUserByEmail ...
func GetUserByEmail(email string) (user *entity.User, has bool, er error) {
	user = &entity.User{Email: email}
	has, er = database.Orm.Table("user").Get(user)
	return
}

// GetUserByIDAndPassword ...
func GetUserByIDAndPassword(userid int64, password string) (user *entity.User, has bool, er error) {
	user = &entity.User{}
	has, er = database.Orm.Where("id=? and password=?", userid, password).Get(user)
	return
}

// GetUserByEmailAndPassword ...
func GetUserByEmailAndPassword(email, password string) (user *entity.User, has bool, er error) {
	user = &entity.User{}
	has, er = database.Orm.Where("email=? and password=?", email, password).Get(user)
	return
}

// CheckUserByEmail ...
func CheckUserByEmail(email string) (has bool, er error) {
	has, er = database.Orm.Where("email=?", email).Get(&entity.User{})
	return
}

// NewUser ...
func NewUser(email, password string) (newUser *entity.User, er error) {
	newUser = &entity.User{
		Email:    email,
		Password: password,
	}
	_, er = database.Orm.Table("user").Insert(newUser)
	return
}

// DeleteUser ...
func DeleteUser(userid int64, password string) (er error) {
	delUser := &entity.User{
		ID:       userid,
		Password: password,
	}
	_, er = database.Orm.Table("user").Delete(delUser)
	return
}

// ChangePassword ...
func ChangePassword(userid int64, newPassword string) (er error) {
	_, er = database.Orm.Table("user").Id(userid).Update(&entity.User{Password: newPassword})
	return
}
