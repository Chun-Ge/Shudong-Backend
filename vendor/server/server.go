package server

import (
	"route"

	"github.com/kataras/iris"
)

// StartServer .
func StartServer() {
	app := iris.New()
	route.RegisterUserRoute(app)
	app.Run(iris.Addr(":8080"))
}
