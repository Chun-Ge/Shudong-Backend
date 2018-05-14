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
	route.Register(app)
	middlewares.Register(app)
	app.Run(iris.Addr("" + args.Port))
}
