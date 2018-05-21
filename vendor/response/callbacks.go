package response

import (
	"github.com/kataras/iris"
)

// GenCallbackInternalServerError generates a callback func
// which calls InternalServerError() by passing null to data
func GenCallbackInternalServerError(ctx iris.Context) func() {
	return func() {
		InternalServerError(ctx, iris.Map{})
	}
}
