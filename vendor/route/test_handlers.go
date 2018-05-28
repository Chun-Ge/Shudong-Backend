package route

import "github.com/kataras/iris"

func registerTestHandler(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>hi, I just exist in order to see if the server is still running</h1>")
	})

	testRoute := app.Party("/test", func(ctx iris.Context) {
		ctx.Writef("test route:")
		ctx.Writef(ctx.Path())
	})

	// in case the auto path correction is disabled
	testRoute.Any("/", func(ctx iris.Context) {
		ctx.Writef("test route ROOT")
	})
	// common test subpath
	testRoute.Any("/{subpath:string}", func(ctx iris.Context) {
		ctx.Writef(ctx.Params().Get("subpath"))
	})
}
