package middlewares

import (
	"github.com/kataras/iris"
)

// Register ...
func Register(app *iris.Application) {
	// register JSON data format check
	// and set Response Content-Type = "application/json"
	registerJSONCheck(app)
	registerXDomainSupport(app)

	registerJwt(app)
}

func registerJwt(app *iris.Application) {
	// Use customized serve for the middleware's action in Iris application.
	// Check and parse JWT at the very beginning of each route,
	// where the ContextKey field and "errjwt" field will be set.
	app.Use(ServeJwt)
}

func registerJSONCheck(app *iris.Application) {
	app.UseGlobal(CheckContentTypeJSON)
	app.UseGlobal(SetResponseContentTypeJSON)
}

func registerXDomainSupport(app *iris.Application) {
	app.UseGlobal(SetAccessControlAllowOrigin)
}
