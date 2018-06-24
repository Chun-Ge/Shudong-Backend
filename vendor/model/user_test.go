package model

import (
	"database"
	"testing"

	"github.com/dmgk/faker"
)

func Test_GetValidUserByID(t *testing.T) {
	database.Start()
	len, e := database.Orm.Table("user").Count();
	if e != nil {
		t.Error("Get valid user by ID failed.")
		t.FailNow()
	}
	var i int64
	for i = 1; i <= len; i++ {
		_, has, e := GetUserByID(i)
		if (e != nil || !has) {
			t.Error("Get valid user by ID failed.")
			t.FailNow()
		}
	}
	t.Log("Get valid user by ID passed.")
}

func Test_GetInValidUserByID_1(t *testing.T) {
	database.Start()
	len, e := database.Orm.Table("user").Count()
	if e != nil {
		t.Error("Get invalid user by ID failed.")
		t.FailNow()
	}
	var i int64
	// TODO: i=0的时候有问题
	for i = -len; i <= 0; i++ {
		user, has, e := GetUserByID(i)
		if (e == nil && has) {
			t.Error("Get invalid user by ID failed.")
			t.Error("The information of user is:", user)
			t.FailNow()
		}
	}
	t.Log("Get invalid user by ID passed.")
}

func Test_GetInValidUserByID_2(t *testing.T) {
	database.Start()
	len, e := database.Orm.Table("user").Count()
	if e != nil {
		t.Error("Get invalid user by ID failed.")
		t.FailNow()
	}
	var i int64
	for i = len + 1; i <= 2 * len; i++ {
		user, has, e := GetUserByID(i)
		if (e == nil && has) {
			t.Error("Get invalid user by ID failed.")
			t.Error("The information of user is:", user)
			t.FailNow()
		}
	}
	t.Log("Get invalid user by ID passed.")
}

func Test_CheckCertainUser(t *testing.T) {
	database.Start()
	len, e := database.Orm.Table("user").Count()
	if e != nil {
		t.Error("Get invalid user by ID failed.")
		t.FailNow()
	}
	email := faker.Internet().Email()
	password := faker.Internet().Password(8, 14)
	user, e := NewUser(email, password)
	if e != nil {
		t.Error("Check certain user failed: create user failed.")
		t.FailNow()
	} else {
		t.Log("Create user passed.")
	}
	// tset GetUserByID
	if userValidate, has, e := GetUserByID(len+1); e != nil || !has || *userValidate != *user {
		t.Error("Check certain user failed: get user by id failed.")
	} else {
		t.Log("Get user by id passed.")
	}
	// test GetUserByEmail
	if userValidate, has, e := GetUserByEmail(email); e != nil || !has || *userValidate != *user {
		t.Error("Check certain user failed: get user by email failed.")
	} else {
		t.Log("Get user by email passed.")
	}
	// test GetUserByIDAndPassword
	if userValidate, has, e := GetUserByIDAndPassword(len+1, password); e != nil || !has || *userValidate != *user {
		t.Error("Check certain user failed: get user by id and password failed.")
	} else {
		t.Log("Get user by id and password passed.")
	}
	// test GetUserByEmailAndPassword
	if userValidate, has, e := GetUserByEmailAndPassword(email, password); e != nil || !has || *userValidate != *user {
		t.Error("Check certain user failed: get user by email and password failed.")
	} else {
		t.Log("Get user by email and password passed.")
	}
	// test CheckUserByEmail
	if has, e := CheckUserByEmail(email); e != nil || !has {
		t.Error("Check certain user failed: check user by email failed.")
	} else {
		t.Log("Check user by email passed.")
	}
	// test ChangePassword
	newPassword := faker.Internet().Password(8, 14)
	if e := ChangePassword(len+1, newPassword); e != nil {
		t.Error("Check certain user failed: change password failed.")
	} else {
		t.Log("Change password passed.")
	}
	// check the correctness of ChangePassword
	if userValidate, has, e := GetUserByEmail(email); e != nil || !has || userValidate.Password != newPassword {
		t.Error("Check certain user failed: the password has not been changed.")
	} else {
		t.Log("The password has been changed.")
	}
}