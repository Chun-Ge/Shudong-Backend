package route

import (
	"middlewares"
	"service"

	"github.com/kataras/iris"
)

// Register ..
func Register(app *iris.Application) {
	registerUserRoutes(app)
	registerPostRoutes(app)
	registerCommentRoutes(app)
	registerUserUpvotePost(app)
	registerUserUpvoteComment(app)
}

func registerUserRoutes(app *iris.Application) {
	app.Post("/login", service.UserLogin)
	app.Post("/logout", middlewares.CheckLoginStatus, service.UserLogout)
	app.Post("/register", service.UserRegister)
}

func registerPostRoutes(app *iris.Application) {
	postRoutes := app.Party("/posts")
	postRoutes.Use(middlewares.CheckLoginStatus)

	// add any subpath below
	//postRoutes.Get("/", service.GetPosts)
	// postRoutes.Get("/{postid:int min(1)}", service.GetPostByID)
	// postRoutes.Get("/{postid:int min(1)}")
	postRoutes.Post("/", service.CreatePost)
	postRoutes.Post("/{postid:int min(1)}", service.CreateComment)

	postRoutes.Post("/{postid:int min(1)}", service.CreateReportPost)
}

func registerCommentRoutes(app *iris.Application) {
	// redundant API "/comments" for "/posts/{postid:int min(1)}/comments"
	commentRoutes := app.Party("/comments")
	commentRoutes.Use(middlewares.CheckLoginStatus)

	// add any subpath below
	// commentRoutes.Get("/", service.GetComments)
}

func registerUserUpvotePost(app *iris.Application) {
	app.Get("/posts/{postid:int min(1)}/like",
		middlewares.CheckLoginStatus, service.UpvotePost)
}

func registerUserUpvoteComment(app *iris.Application) {
	app.Get("/posts/{postid:int min(1)}/comments/{commentid:int min(1)}/like",
		middlewares.CheckLoginStatus, service.UpvoteComment)
}
