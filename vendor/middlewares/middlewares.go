package middlewares

import (
	"github.com/kataras/iris"
)

// Register .
func Register(app *iris.Application) {
	registerJwt(app)
}

func registerJwt(app *iris.Application) {
	// Use customized serve for the middleware's action in Iris application.
	// Check and parse JWT at the very beginning of each route,
	// where the ContextKey field and "errjwt" field will be set.
	app.Use(ServeJwt)
}
