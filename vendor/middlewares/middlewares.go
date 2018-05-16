package middlewares

import (
	"github.com/kataras/iris"
)

// Register .
func Register(app *iris.Application) {
	registerJwt(app)
}

func registerJWT(app *iris.Application) {
	// Use customized serve for the middleware's action in Iris application.
	app.Use(ServeJWT)
}
