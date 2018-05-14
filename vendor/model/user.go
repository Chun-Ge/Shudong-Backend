package model

import (
	"database"
	"entity"
	"err"
)

// GetUser .
func GetUser(username string, password string) (user *entity.User, er error) {
	_, er = database.Orm.Where("username=? and password=?", username, password).Get(&user)
	err.CheckErr(er)
	return
}
