package model

import (
	"database"
	"entity"
	"testing"

	"github.com/dmgk/faker"
)

func init() {
	database.Start()
}

func Test_GetValidUserByID(t *testing.T) {
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

// 要额外写个User间的比较函数，而不能直接用*a==*b来判断
// 因为下面通过NewUser返回的user的时间戳与数据库里面的时间戳
// 可能存在一定偏差，所以所有时间戳不能成为判断user相等的依据
func userEqual(a *entity.User, b *entity.User) bool {
	return a.ID == b.ID && a.Email == b.Email && a.Password == b.Password
}

func Test_CheckCertainUser(t *testing.T) {
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
	// test GetUserByID
	if userValidate, has, e := GetUserByID(len+1); e != nil || !has || !userEqual(user, userValidate) {
		t.Error("Check certain user failed: get user by id failed.")
		t.Error("Result:", userValidate)
		t.Error("Expect:", user)
	} else {
		t.Log("Get user by id passed.")
	}
	// test GetUserByEmail
	if userValidate, has, e := GetUserByEmail(email); e != nil || !has || !userEqual(user, userValidate) {
		t.Error("Check certain user failed: get user by email failed.")
		t.Error("Result:", userValidate);
		t.Error("Expect:", user)
	} else {
		t.Log("Get user by email passed.")
	}
	// test GetUserByIDAndPassword
	if userValidate, has, e := GetUserByIDAndPassword(len+1, password); e != nil || !has || !userEqual(user, userValidate) {
		t.Error("Check certain user failed: get user by id and password failed.")
		t.Error("Result:", userValidate)
		t.Error("Expect:", user)
	} else {
		t.Log("Get user by id and password passed.")
	}
	// test GetUserByEmailAndPassword
	if userValidate, has, e := GetUserByEmailAndPassword(email, password); e != nil || !has || !userEqual(user, userValidate) {
		t.Error("Check certain user failed: get user by email and password failed.")
		t.Error("Result:", userValidate)
		t.Error("Expect:", user)
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
		t.Error("Result:", userValidate)
		t.Error("Expect:", user)
	} else {
		t.Log("The password has been changed.")
	}
}