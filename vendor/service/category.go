package service

import (
	"err"
	"model"
	"response"

	"github.com/kataras/iris"
)

// GetAllCategoryNames ...
func GetAllCategoryNames(ctx iris.Context) {
	categoryNames, er := model.GetAllCategoryNames()
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{
		"categoryNames": categoryNames,
	})
}
