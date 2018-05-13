package route

import (
	"service"

	"github.com/kataras/iris"
)

// Register ..
func Register(app *iris.Application) {
	registerUserRoutes(app)
	registerPostRoutes(app)
	registerCommentRoutes(app)
}

func registerUserRoutes(app *iris.Application) {
	app.Post("/login", service.UserLogin)
}

func registerPostRoutes(app *iris.Application) {

}

func registerCommentRoutes(app *iris.Application) {

}
