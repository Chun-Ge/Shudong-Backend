package model

import (
	_ "errors"

	"database"
	"entity"
	_ "err"
)

// check auth_code by user_id and code
func CheckAuthCodeByUserAndCode(userID int64, code string) (bool, error) {
	return database.Orm.Table("auth_code").Exist(
		&entity.AuthCode{
			UserID: userID,
			Code:   code,
		})
}

// GetAuthCodeByUserAndCode
func GetAuthCodeByUserAndCode(userID int64, code string) (*entity.AuthCode, bool, error) {
	ret := &entity.AuthCode{
		UserID: userID,
		Code:   code,
	}
	has, er := database.Orm.Table("auth_code").Get(ret)
	return ret, has, er
}

// insert an auth code
func NewAuthCode(userID int64, code string) (ret *entity.AuthCode, er error) {
	ret = &entity.AuthCode{UserID: userID, Code: code}
	_, er = database.Orm.Table("auth_code").Insert(ret)
	return
}

// CheckAuthByUser
func CheckAuthCodeByUser(userID int64) (bool, error) {
	return database.Orm.Table("auth_code").Exist(
		&entity.AuthCode{
			UserID: userID,
		})
}

// update the auth code
func UpdateAuthCode(userID int64, code string) (er error) {
	_, er = database.Orm.Table("auth_code").Where("user_id=?", userID).Update(
		map[string]interface{}{"code": code})
	return
}
