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

	registerRootRoutes(app)
	registerUserRoutes(app)
	registerPostRoutes(app)
	registerCommentRoutes(app)
}

// registerRootRoutes .
func registerRootRoutes(app *iris.Application) {
	app.Post("/login", service.UserLogin)
	app.Post("/logout", middlewares.CheckLoginStatus, service.UserLogout)
	app.Post("/reset_password/authcode", service.GenAuthCode)
	app.Patch("/reset_password", service.ResetPassword)
}

// registerUserRoutes .
// Group url of "/users"
func registerUserRoutes(app *iris.Application) {
	app.Post("/users", service.UserRegister)

	userRoutes := app.Party("/users")
	userRoutes.Use(middlewares.CheckLoginStatus)

	userRoutes.Patch("/password", service.ChangePassword)
}

// registerPostRoutes .
// Group url of "/posts"
func registerPostRoutes(app *iris.Application) {
	postRoutes := app.Party("/posts")
	postRoutes.Use(middlewares.CheckLoginStatus)

	// add any subpath below
	// postRoutes.Get("/", service.GetPosts)
	// postRoutes.Get("/{postid:int min(1)}", service.GetPostByID)
	postRoutes.Post("/", service.CreatePost)
	postRoutes.Delete("/", service.DeletePost)

	postRoutes.Get("/{postid:int min(1)}/like", service.UpvotePost)

	postRoutes.Post("/{postid:int min(1)}/report", service.CreateReportPost)

	postRoutes.Get("/{postid:int min(1)}/star", service.StarPost)
}

// registerCommentRoutes .
// Group url of "/post/{postid}/comments"
func registerCommentRoutes(app *iris.Application) {
	// redundant API "/comments" for "/posts/{postid:int min(1)}/comments"
	commentRoutes := app.Party("/posts/{postid:int min(1)}/comments")
	commentRoutes.Use(middlewares.CheckLoginStatus)

	// add any subpath below
	// app.Get("/", service.GetComments)
	commentRoutes.Post("/", service.CreateComment)
	commentRoutes.Delete("/", service.DeleteComment)

	commentRoutes.Get("/{commentid:int min(1)}/like", service.UpvoteComment)

	commentRoutes.Post("/{commentid:int min(1)}/report", service.CreateReportComment)
}
