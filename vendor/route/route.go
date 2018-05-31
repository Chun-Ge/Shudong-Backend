package route

import (
	"args"
	"middlewares"
	"service"

	"github.com/kataras/iris"
)

// Register ..
func Register(app *iris.Application) {
	if args.DEBUG {
		registerTestHandler(app)
	}

	registerUserRoutes(app)
	registerPostRoutes(app)
	registerCommentRoutes(app)
	registerUserUpvotePost(app)
	registerUserUpvoteComment(app)
	registerReportPost(app)
	registerReportComment(app)
	registerUserStarPost(app)
}

func registerUserRoutes(app *iris.Application) {
	app.Post("/login", service.UserLogin)
	app.Post("/logout", middlewares.CheckLoginStatus, service.UserLogout)
	app.Post("/register", service.UserRegister)
	app.Put("/users/change_password", service.ChangePassword)
	// app.Post("/users/reset_password/authcode", service.GenAuthCode)
	// app.Put("/users/reset_password", service.ResetPassword)
}

func registerPostRoutes(app *iris.Application) {
	// postRoutes := app.Party("/posts")
	// postRoutes.Use(middlewares.CheckLoginStatus)

	// add any subpath below
	// postRoutes.Get("/", service.GetPosts)
	// postRoutes.Get("/{postid:int min(1)}", service.GetPostByID)

	// app.Get("/posts", middlewares.CheckLoginStatus, service.GetPosts)
	// app.Get("/{postid:int min(1)}", middlewares.CheckLoginStatus, service.GetPostByID)
	app.Post("/posts", middlewares.CheckLoginStatus, service.CreatePost)
	// postRoutes.Delete("/{postid:int min(1)}", service.DeletePost)
}

func registerCommentRoutes(app *iris.Application) {
	// redundant API "/comments" for "/posts/{postid:int min(1)}/comments"
	// commentRoutes := app.Party("/post/{postid:int min(1)}").Party("/comments")
	// commentRoutes.Use(middlewares.CheckLoginStatus)

	// add any subpath below
	// app.Get("/", service.GetComments)
	app.Post("/posts/{postid:int min(1)}", middlewares.CheckLoginStatus, service.CreateComment)
	app.Delete("/{commentid:int min(1)}", middlewares.CheckLoginStatus, service.DeleteComment)
}

func registerUserUpvotePost(app *iris.Application) {
	app.Get("/posts/{postid:int min(1)}/like",
		middlewares.CheckLoginStatus, service.UpvotePost)
}

func registerUserUpvoteComment(app *iris.Application) {
	app.Get("/posts/{postid:int min(1)}/comments/{commentid:int min(1)}/like",
		middlewares.CheckLoginStatus, service.UpvoteComment)
}

func registerReportPost(app *iris.Application) {
	app.Post("/posts/{postid:int min(1)}/report",
		middlewares.CheckLoginStatus, service.CreateReportPost)
}

func registerReportComment(app *iris.Application) {
	app.Post("/posts/{postid:int min(1)}/comments/{commentid:int min(1)}/report",
		middlewares.CheckLoginStatus, service.CreateReportComment)
}

func registerUserStarPost(app *iris.Application) {
	app.Get("/posts/{postid:int min(1)}/star",
		middlewares.CheckLoginStatus, service.StarPost)
}
