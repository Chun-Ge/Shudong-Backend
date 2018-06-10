package route

import (
	"response"

	"github.com/kataras/iris"
)

func registerTestHandler(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>hi, I just exist in order to see if the server is still running</h1>")
	})

	testRoute := app.Party("/test", func(ctx iris.Context) {
		// Add some log codes.

		ctx.Next()
	})

	// in case the auto path correction is disabled
	testRoute.Any("/", func(ctx iris.Context) {
		response.OK(ctx, iris.Map{
			"test-route-root": "See this when hitting /test/",
		})
	})
	// common test subpath
	testRoute.Any("/{subpath:string}", func(ctx iris.Context) {
		response.OK(ctx, iris.Map{
			"test-route":   ctx.Path(),
			"test-subpath": ctx.Params().Get("subpath"),
		})
	})
}
