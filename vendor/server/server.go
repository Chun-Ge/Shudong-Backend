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

	middlewares.Register(app)
	route.Register(app)

	app.Run(iris.Addr("" + args.Port))
}

// StartWithConfiguration reads the config file and
func StartWithConfiguration(configFilePath string) {
	app := iris.New()

	middlewares.Register(app)
	route.Register(app)

	app.Configure(iris.WithConfiguration(iris.YAML(configFilePath)))

	app.Run(iris.Addr("" + args.Port))
}
