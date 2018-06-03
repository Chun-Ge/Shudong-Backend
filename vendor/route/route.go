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

	// User Operation url
	// Retrieve user info
	// userRoutes.Get("{userId:int min(1)}", handler)
	// userRoutes.PUT("{userId:int min(1)}", handler)
	// userRoutes.PATCH("{userId:int min(1)}", handler)
	// userRoutes.Delete("{userId:int min(1)}", handler)
}

// registerPostRoutes .
// Group url of "/posts"
func registerPostRoutes(app *iris.Application) {
	postRoutes := app.Party("/posts")
	postRoutes.Use(middlewares.CheckLoginStatus)

	// subpath of "/posts"
	// Post Collection and Creation
	// postRoutes.Get("/", handler)
	postRoutes.Post("/", service.CreatePost)

	// Get and Delete Post
	// postRoutes.Get("/{postId:int min(1)", handler)
	postRoutes.Delete("/{postId:int min(1)", service.DeletePost)

	// share a post
	// postRoutes.Get("/{postId:int min(1)/share", handler)

	// liek/un-like a post
	postRoutes.Get("/{postId:int min(1)}/like", service.UpvotePost)

	// report a post
	postRoutes.Post("/{postId:int min(1)}/report", service.CreateReportPost)

	// star a post
	postRoutes.Get("/{postId:int min(1)}/star", service.StarPost)
}

// registerCommentRoutes .
// Group url of "/post/{postId}/comments"
func registerCommentRoutes(app *iris.Application) {
	// redundant API "/comments" for "/posts/{postId:int min(1)}/comments"
	commentRoutes := app.Party("/posts/{postId:int min(1)}/comments")
	commentRoutes.Use(middlewares.CheckLoginStatus)

	// add any subpath of "/posts/{postId:int min(1)}/comments"
	// Comment Collection and Creation
	// commentRoutes.Get("/", handler)
	commentRoutes.Post("/", service.CreateComment)

	// delete comment
	commentRoutes.Delete("/{commentId:int min(1)", service.DeleteComment)

	// like/un-like a comment
	commentRoutes.Get("/{commentId:int min(1)}/like", service.UpvoteComment)

	// resport a comment
	commentRoutes.Post("/{commentId:int min(1)}/report", service.CreateReportComment)
}
