package response

import (
	"github.com/kataras/iris"
)

// GenCallbackInternalServerError generates a callback func
// which calls InternalServerError() by passing null to data
func GenCallbackInternalServerError(ctx iris.Context, err error) func() {
	return func() {
		InternalServerError(ctx, iris.Map{})
		ctx.StopExecution()
		panic(err)
	}
}

// GenCallbackBadRequest generates a callback func
// which calls BadRequest() by passing null to data
func GenCallbackBadRequest(ctx iris.Context, err error) func() {
	return func() {
		BadRequest(ctx, iris.Map{})
		ctx.StopExecution()
		panic(err)
	}
}
