package model

import (
	"database"
	"entity"
	"err"
)

// GetUserByID .
func GetUserByID(userid int64) (*entity.User, error) {
	ret := &entity.User{ID: userid}
	_, er := database.Orm.Table("user").Get(ret)
	return ret, er
}

// GetUserByEmailAndPassword .
func GetUserByEmailAndPassword(email, password string) (user *entity.User, er error) {
	_, er = database.Orm.Where("email=? and password=?", email, password).Get(&user)
	err.CheckErr(er)
	return
}

// NewUser .
func NewUser(email, password string) (er error) {
	newUser := &entity.User{
		Email:    email,
		Password: password,
	}
	_, er = database.Orm.Table("user").Insert(newUser)
	return
}

// DeleteUser .
func DeleteUser(userid int64, password string) (er error) {
	delUser := &entity.User{
		ID:       userid,
		Password: password,
	}
	_, er = database.Orm.Table("user").Delete(delUser)
	return
}

// ChangePassword .
func ChangePassword(userid int64, newPassword string) (er error) {
	_, er = database.Orm.Table("user").Id(userid).Update(&entity.User{Password: newPassword})
	return
}
