package service

import (
	"args"
	"crypto/md5"
	"encoding/hex"
	"entity"
	"err"
	"model"
	"response"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

// UserFormData .
type UserFormData struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func encodePassword(initPassword string) (password string) {
	md5Hash := md5.New()
	md5Hash.Write([]byte(initPassword))
	password = hex.EncodeToString(md5Hash.Sum(nil))
	return
}

// UserLogin .
func UserLogin(ctx iris.Context) {
	var user entity.User
	var has bool
	userForm := UserFormData{}

	er := ctx.ReadForm(&userForm)
	err.CheckErrWithPanic(er)

	email := userForm.Email
	password := encodePassword(userForm.Password)

	user, has, er = model.GetUserByEmailAndPassword(email, password)
	err.CheckErrWithPanic(er)
	if !has {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24).Unix()},
	})
	t, er := token.SignedString([]byte(args.SecretKey))
	err.CheckErrWithPanic(er)

	ctx.ResponseWriter().Header().Set("Authorization", "Bearer "+t)
	response.OK(ctx, iris.Map{
		"userid": user.ID,
	})
}

// UserLogout .
func UserLogout(ctx iris.Context) {
	delete(ctx.ResponseWriter().Header(), "Authorization")
	response.OK(ctx, iris.Map{})
}

// UserRegister .
func UserRegister(ctx iris.Context) {
	userForm := UserFormData{}

	ctx.ReadForm(&userForm)

	email := userForm.Email
	password := encodePassword(userForm.Password)

	has, er := model.CheckUserByEmail(email)
	err.CheckErrWithPanic(er)
	if has {
		response.Conflict(ctx, iris.Map{})
		return
	}

	user, er := model.NewUser(email, password)
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{
		"userid": user.ID,
	})
}
