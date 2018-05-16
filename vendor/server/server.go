package server

import (
	"args"
	"middlewares"
	"route"

	"github.com/kataras/iris"
)

// StartServer .
func StartServer() {
	app := iris.New()
	middlewares.Register(app)
	route.Register(app)
	app.Run(iris.Addr("" + args.Port))
}
