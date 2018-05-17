package service

import (
	"args"
	"crypto/md5"
	"encoding/hex"
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
	userForm := UserFormData{}

	if err := ctx.ReadForm(&userForm); err != nil {
		response.Unauthorized(ctx, iris.Map{
			"msg":  "Unauthorized",
			"data": iris.Map{},
		})
		return
	}

	email := userForm.Email
	password := encodePassword(userForm.Password)

	user, has, err := model.GetUserByEmailAndPassword(email, password)
	if err != nil || !has {
		response.Unauthorized(ctx, iris.Map{
			"msg":  "Unauthorized",
			"data": iris.Map{},
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24).Unix()},
	})
	t, _ := token.SignedString([]byte(args.SecretKey))

	ctx.ResponseWriter().Header().Set("Authorization", "Bearer "+t)
	response.OK(ctx, iris.Map{
		"msg": "OK",
		"data": iris.Map{
			"userid": user.ID,
		}})
}

// UserRegister .
func UserRegister(ctx iris.Context) {
	userForm := UserFormData{}

	if err := ctx.ReadForm(&userForm); err != nil {
		response.InternalServerError(ctx, iris.Map{
			"msg":  "Internal Server Error",
			"data": iris.Map{},
		})
		return
	}

	email := userForm.Email
	password := encodePassword(userForm.Password)

	has, err := model.CheckUserByEmail(email)
	if err != nil {
		response.InternalServerError(ctx, iris.Map{
			"msg":  "Internal Server Error",
			"data": iris.Map{},
		})
		return
	}
	if has {
		response.Conflict(ctx, iris.Map{
			"msg":  "Conflict",
			"data": iris.Map{},
		})
		return
	}

	user, err := model.NewUser(email, password)
	if err != nil {
		response.InternalServerError(ctx, iris.Map{
			"msg":  "Internal Server Error",
			"data": iris.Map{},
		})
		return
	}

	response.OK(ctx, iris.Map{
		"msg": "OK",
		"data": iris.Map{
			"userid": user.ID,
		}})
}
