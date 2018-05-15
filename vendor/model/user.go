package model

import (
	"database"
	"entity"
	"err"
)

// GetUser .
func GetUser(email, password string) (user *entity.User, er error) {
	_, er = database.Orm.Where("email=? and password=?", email, password).Get(&user)
	err.CheckErr(er)
	return
}
