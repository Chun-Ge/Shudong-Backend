package service

import (
	"crypto/md5"
	"database"
	"encoding/hex"
	"entity"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

// UserFormData .
type UserFormData struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func encodePassword(initPassword string) (password string) {
	md5Hash := md5.New()
	md5Hash.Write([]byte(initPassword))
	password = hex.EncodeToString(md5Hash.Sum(nil))
	return
}

// UserLogin ..
func UserLogin(ctx iris.Context) {
	userForm := UserFormData{}

	if err := ctx.ReadForm(&userForm); err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		return
	}

	username := userForm.Username
	password := encodePassword(userForm.Password)

	user := new(entity.User)
	if _, err := database.Orm.Where("username=? and password=?", username, password).Get(&user); err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.ID,
		"timeout": jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24).Unix()},
	})
	t, _ := token.SignedString([]byte(args.SecretKey))

	ctx.SetCookieKV("Token", t)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"msg": "OK",
		"data": iris.Map{
			"username": username,
		},
	})
}
