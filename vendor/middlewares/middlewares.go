package middlewares

import (
	"args"

	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

// Register .
func Register(app *iris.Application) {
	registerJwt(app)
}

func registerJwt(app *iris.Application) {
	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(args.Se), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	app.Use(myJwtMiddleware.Serve)
}
