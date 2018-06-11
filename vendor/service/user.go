package service

import (
	"args"
	"crypto/md5"
	"encoding/hex"
	"entity"
	"err"
	"fmt"
	"math"
	"math/rand"
	"middlewares"
	"model"
	"response"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

// UserRequestData ...
type UserRequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ChangePasswordRequestData ...
type ChangePasswordRequestData struct {
	UserID      int64
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

// GenAuthCodeRequestData ...
type GenAuthCodeRequestData struct {
	Email string `json:"email"`
}

// ResetPasswordRequestData ...
type ResetPasswordRequestData struct {
	Email       string `json:"email"`
	AuthCode    string `json:"authCode"`
	NewPassword string `json:"newPassword"`
}

func encodePassword(initPassword string) (password string) {
	md5Hash := md5.New()
	md5Hash.Write([]byte(initPassword))
	password = hex.EncodeToString(md5Hash.Sum(nil))
	return
}

// UserLogin ...
func UserLogin(ctx iris.Context) {
	user := &entity.User{}
	var has bool
	userRequest := &UserRequestData{}

	er := ctx.ReadJSON(userRequest)
	err.CheckErrWithCallback(er, response.GenCallbackBadRequest(ctx, er))

	email := userRequest.Email
	password := encodePassword(userRequest.Password)

	user, has, er = model.GetUserByEmailAndPassword(email, password)
	err.CheckErrWithPanic(er)
	if !has {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	t, er := token.SignedString([]byte(args.SecretKey))
	err.CheckErrWithPanic(er)

	ctx.ResponseWriter().Header().Set("Access-Control-Expose-Headers", "Authorization")
	ctx.ResponseWriter().Header().Set("Authorization", "Bearer "+t)
	response.OK(ctx, iris.Map{
		"userId": user.ID,
	})
}

// UserLogout ...
func UserLogout(ctx iris.Context) {
	delete(ctx.ResponseWriter().Header(), "Authorization")
	response.OK(ctx, iris.Map{})
}

// UserRegister ...
func UserRegister(ctx iris.Context) {
	userRequest := &UserRequestData{}

	er := ctx.ReadJSON(userRequest)
	err.CheckErrWithCallback(er, response.GenCallbackBadRequest(ctx, er))

	email := userRequest.Email
	password := encodePassword(userRequest.Password)

	has, er := model.CheckUserByEmail(email)
	err.CheckErrWithPanic(er)
	if has {
		response.Conflict(ctx, iris.Map{})
		return
	}

	user, er := model.NewUser(email, password)
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{
		"userId": user.ID,
	})
}

// ChangePassword ...
// route: [/users/change_password] [PUT]
// pre: the user is in the session
// post: the password has been updated
func ChangePassword(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)

	changePasswordRequest := &ChangePasswordRequestData{UserID: userID}

	er := ctx.ReadJSON(changePasswordRequest)
	err.CheckErrWithCallback(er, response.GenCallbackBadRequest(ctx, er))

	oldPassword := encodePassword(changePasswordRequest.OldPassword)
	newPassword := encodePassword(changePasswordRequest.NewPassword)

	_, has, er := model.GetUserByIDAndPassword(userID, oldPassword)
	if er != nil || !has {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	er = model.ChangePassword(userID, newPassword)
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{})
}

// GenAuthCode for reset password
// route: [/users/reset_password/authcode] [POST]
// pre: None
// post: store the map info of auth code of the user
func GenAuthCode(ctx iris.Context) {
	genAuthCodeRequest := &GenAuthCodeRequestData{}

	er := ctx.ReadJSON(genAuthCodeRequest)
	err.CheckErrWithCallback(er, response.GenCallbackBadRequest(ctx, er))

	email := genAuthCodeRequest.Email

	user, has, er := model.GetUserByEmail(email)
	err.CheckErrWithPanic(er)

	if !has {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	has, er = model.CheckAuthCodeByUser(user.ID)
	err.CheckErrWithPanic(er)

	newCode := genRandAuthCode(args.AuthCodeSize)
	// if exists, update, or, insert
	if !has {
		_, er := model.NewAuthCode(user.ID, newCode)
		err.CheckErrWithPanic(er)
	} else {
		er = model.UpdateAuthCode(user.ID, newCode)
		err.CheckErrWithPanic(er)
	}

	response.OK(ctx, iris.Map{
		"authCode": newCode,
	})
}

// ResetPassword ...
// route : [/users/reset_password] [PUT]
// pre: there are 3 key in the request JSON: "email", "authCode", "newPassword"
// post: if authCode is valid with the user, the password will have been reset
func ResetPassword(ctx iris.Context) {
	info := &ResetPasswordRequestData{}

	er := ctx.ReadJSON(info)
	err.CheckErrWithCallback(er, response.GenCallbackBadRequest(ctx, er))

	// Check whether the email is valid.
	user, has, er := model.GetUserByEmail(info.Email)
	err.CheckErrWithPanic(er)
	if !has {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	// Check whether the code is stored in the database.
	authCode, has, er := model.GetAuthCodeByUserAndCode(user.ID, info.AuthCode)
	err.CheckErrWithPanic(er)
	if !has {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	yes, er := isBefore(args.AuthCodeLifeTime, authCode.UpdateTime)
	err.CheckErrWithPanic(er)

	// now - AuthCodeLifeTime(minutes) is not before codeUpdateTime
	if !yes {
		// Destroy the authCode if outdated.
		_, er := model.DeleteAuthCode(authCode.ID)
		err.CheckErrWithPanic(er)

		response.Forbidden(ctx, iris.Map{})
		return
	}

	er = model.ChangePassword(user.ID, encodePassword(info.NewPassword))
	err.CheckErrWithPanic(er)

	// Destroy the authCode if the password is successfully changed.
	_, er = model.DeleteAuthCode(authCode.ID)
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{})
}

// gen numeric auth code with size bits
func genRandAuthCode(size int) string {
	maxOne := int32(math.Pow10(size))
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	// return string(rnd.Int31n(maxOne))
	// return "98765"
	strFormat := "%0" + strconv.Itoa(size) + "d"
	return fmt.Sprintf(strFormat, rnd.Int31n(maxOne))
}

// now() - lefiTime(minutes) is before pastTime
func isBefore(lifeTime int, pastTime time.Time) (bool, error) {
	now := time.Now()
	// m, er := time.ParseDuration("-" + string(lifeTime) + "m")
	durationParam := "-" + strconv.Itoa(lifeTime) + "m"
	m, er := time.ParseDuration(durationParam)

	// AuthCodeLifeTime minutes ago
	cmpTime := now.Add(m)
	return cmpTime.Before(pastTime), er
}
