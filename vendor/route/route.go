package route

import (
	"args"
	"middlewares"
	"service"

	"github.com/kataras/iris"
)

// Register ...
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
	app.Post("/login", service.UserLogin).Name = "UserLogin"
	app.Post("/logout", middlewares.CheckLoginStatus, service.UserLogout).Name = "UserLogout"
	app.Post("/reset_password/authcode", service.GenAuthCode).Name = "GenAuthCode"
	app.Patch("/reset_password", service.ResetPassword).Name = "ResetPassword"
}

// registerUserRoutes .
// Group url of "/users"
func registerUserRoutes(app *iris.Application) {
	app.Post("/users", service.UserRegister).Name = "UserRegister"

	userRoutes := app.Party("/users")
	userRoutes.Use(middlewares.CheckLoginStatus)

	userRoutes.Patch("/password", service.ChangePassword).Name = "ChangePassword"
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
	postRoutes.Get("/", service.GetRecentPosts).Name = "GetRecentPosts"
	postRoutes.Post("/", service.CreatePost).Name = "CreatePost"

	// Get and Delete Post
	postRoutes.Get("/{postId:int min(1)}", middlewares.CheckPostIDExistence,
		service.GetPostByID).Name = "GetPostByID"
	postRoutes.Delete("/{postId:int min(1)}", middlewares.CheckPostIDExistence,
		service.DeletePost).Name = "DeletePost"

	// share a post
	// postRoutes.Get("/{postId:int min(1)/share", handler)

	// liek/un-like a post
	postRoutes.Get("/{postId:int min(1)}/like", middlewares.CheckPostIDExistence,
		service.UpvotePost).Name = "UpvotePost"

	// report a post
	postRoutes.Post("/{postId:int min(1)}/report", middlewares.CheckPostIDExistence,
		service.CreateReportPost).Name = "CreateReportPost"

	// star a post
	postRoutes.Get("/{postId:int min(1)}/star", middlewares.CheckPostIDExistence,
		service.StarPost).Name = "StarPost"

	// share a post
	postRoutes.Get("/{postId:int min(1)}/share", middlewares.CheckPostIDExistence,
		service.SharePost).Name = "SharePost"

	// all category names
	postRoutes.Get("/categories", service.GetAllCategoryNames).Name = "GetAllCategoryNames"
}

// registerCommentRoutes .
// Group url of "/post/{postId}/comments"
func registerCommentRoutes(app *iris.Application) {
	// redundant API "/comments" for "/posts/{postId:int min(1)}/comments"
	commentRoutes := app.Party("/posts/{postId:int min(1)}/comments")
	commentRoutes.Use(middlewares.CheckLoginStatus)
	commentRoutes.Use(middlewares.CheckPostIDExistence)

	// add any subpath of "/posts/{postId:int min(1)}/comments"
	// Comment Collection and Creation
	// commentRoutes.Get("/", handler)
	commentRoutes.Post("/", service.CreateComment).Name = "CreateComment"

	// Get comments of a specific post
	commentRoutes.Get("/", service.GetCommentsOfAPost).Name = "GetCommentsOfAPost"

	// delete comment
	commentRoutes.Delete("/{commentId:int min(1)}", middlewares.CheckCommentIDExistenceAndLegitimate,
		service.DeleteComment).Name = "DeleteComment"

	// like/un-like a comment
	commentRoutes.Get("/{commentId:int min(1)}/like", middlewares.CheckCommentIDExistenceAndLegitimate,
		service.UpvoteComment).Name = "UpvoteComment"

	// resport a comment
	commentRoutes.Post("/{commentId:int min(1)}/report", middlewares.CheckCommentIDExistenceAndLegitimate,
		service.CreateReportComment).Name = "CreateReportComment"
}
