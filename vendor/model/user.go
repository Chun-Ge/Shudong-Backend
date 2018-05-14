package model

import (
	"database"
	"entity"
)

// GetUser .
func GetUser(username string, password string) (user *entity.User, err error) {
	_, err = database.Orm.Where("username=? and password=?", username, password).Get(&user)
	checkErr(err)
	return
}
