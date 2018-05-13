package server

import (
	"args"
	"middlewares"
	"route"

	"github.com/kataras/iris"
)

// Start .
func Start() {
	app := iris.New()
	// database.InitTable()

	route.Register(app)
	middlewares.Register(app)
	app.Run(iris.Addr("" + args.Port))
}
