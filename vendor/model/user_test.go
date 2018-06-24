package model

import (
	"database"
	"testing"

	"github.com/dmgk/faker"
)

func Test_GetValidUserByID(t *testing.T) {
	len, e := database.Orm.Table("user").Count();
	if e != nil {
		t.Error("Get valid user by ID failed.")
	}
	var i int64
	for i = 1; i <= len; i++ {
		_, has, e := GetUserByID(i)
		if (e != nil || !has) {
			t.Error("Get valid user by ID failed.")
		}
	}
	t.Log("Get valid user by ID passed.")
}

func Test_GetInValidUserByID_1(t *testing.T) {
	len, e := database.Orm.Table("user").Count()
	if e != nil {
		t.Error("Get invalid user by ID failed.")
	}
	var i int64
	for i = -len; i <= 0; i++ {
		_, has, e := GetUserByID(i)
		if (e != nil || has) {
			t.Error("Get invalid user by ID failed.")
		}
	}
	t.Log("Get invalid user by ID passed.")
}

func Test_GetInValidUserByID_2(t *testing.T) {
	len, e := database.Orm.Table("user").Count()
	if e != nil {
		t.Error("Get invalid user by ID failed.")
	}
	var i int64
	for i = len + 1; i <= 2 * len; i++ {
		_, has, e := GetUserByID(i)
		if (e != nil || has) {
			t.Error("Get invalid user by ID failed.")
		}
	}
	t.Log("Get invalid user by ID passed.")
}

func Test_CheckCertainUser(t *testing.T) {
	len, e := database.Orm.Table("user").Count()
	if e != nil {
		t.Error("Get invalid user by ID failed.")
	}
	email := faker.Internet().Email()
	password := faker.Internet().Password(8, 14)
	user, e := NewUser(email, password)
	if e != nil {
		t.Error("Check certain user failed: create user failed.")
	}
	// tet GetUserByID
	if userValidate, has, e := GetUserByID(len+1); e != nil || !has || *userValidate != *user {
		t.Error("Check certain user failed: get user by id failed.")
	}
	// test GetUserByEmail
	if userValidate, has, e := GetUserByEmail(email); e != nil || !has || *userValidate != *user {
		t.Error("Check certain user failed: get user by email failed.")
	}
	// test GetUserByIDAndPassword
	if userValidate, has, e := GetUserByIDAndPassword(len+1, password); e != nil || !has || *userValidate != *user {
		t.Error("Check certain user failed: get user by id and password failed.")
	}
	// test GetUserByEmailAndPassword
	if userValidate, has, e := GetUserByEmailAndPassword(email, password); e != nil || !has || *userValidate != *user {
		t.Error("Check certain user failed: get user by email and password failed.")
	}
	// test CheckUserByEmail
	if has, e := CheckUserByEmail(email); e != nil || !has {
		t.Error("Check certain user failed: check user by email failed.")
	}
	// test ChangePassword
	newPassword := faker.Internet().Password(8, 14)
	if e := ChangePassword(len+1, newPassword); e != nil {
		t.Error("Check certain user failed: change password failed.")
	}
	// check the correctness of ChangePassword
	if userValidate, has, e := GetUserByEmail(email); e != nil || !has || userValidate.Password != newPassword {
		t.Error("Check certain user failed: get user by email failed.")
	}
}