package model

import (
	"database"
	"entity"
)

// CheckAuthCodeByUserAndCode checks auth_code by user_id and code
func CheckAuthCodeByUserAndCode(userID int64, code string) (bool, error) {
	return database.Orm.Table("auth_code").Exist(
		&entity.AuthCode{
			UserID: userID,
			Code:   code,
		})
}

// GetAuthCodeByUserAndCode ...
func GetAuthCodeByUserAndCode(userID int64, code string) (*entity.AuthCode, bool, error) {
	ret := &entity.AuthCode{
		UserID: userID,
		Code:   code,
	}
	has, er := database.Orm.Table("auth_code").Get(ret)
	return ret, has, er
}

// NewAuthCode creates an auth code and inserts it to the DB.
func NewAuthCode(userID int64, code string) (*entity.AuthCode, error) {
	ret := &entity.AuthCode{
		UserID: userID,
		Code:   code,
	}
	_, er := database.Orm.Table("auth_code").Insert(ret)
	return ret, er
}

// CheckAuthCodeByUser ...
func CheckAuthCodeByUser(userID int64) (bool, error) {
	return database.Orm.Table("auth_code").Exist(
		&entity.AuthCode{
			UserID: userID,
		})
}

// UpdateAuthCode ...
func UpdateAuthCode(userID int64, code string) error {
	_, er := database.Orm.Table("auth_code").Where("user_id=?", userID).Update(
		map[string]interface{}{"code": code})
	return er
}
