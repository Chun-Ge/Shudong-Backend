package model

import (
	"database"
	"entity"
	e "err"
)

// GetUser .
func GetUser(username string, password string) (user *entity.User, err error) {
	_, err = database.Orm.Where("username=? and password=?", username, password).Get(&user)
	e.CheckErr(err)
	return
}
