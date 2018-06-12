package middlewares

import (
	"args"
	"fmt"
	"response"

	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

// MyJwtMiddleware ...
var MyJwtMiddleware *jwtmiddleware.Middleware

func init() {
	MyJwtMiddleware = jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(args.SecretKey), nil
		},
		ContextKey: "jwt",
	})
}

// checkJWT the main functionality, checks for token.
// If error raised when parsing JWT,
// "Unauthorized" will be set to "errjwt" field of values of the context.
// Otherwise, the JWT will be set directly to the current ContextKey field
// (which is "jwt" under this project).
func checkJWT(ctx iris.Context, m *jwtmiddleware.Middleware) (string, error) {
	if !m.Config.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return "OK", nil
		}
	}

	token, err := m.Config.Extractor(ctx)

	if err != nil {
		return "Unauthorized", fmt.Errorf("Error extracting token: %v", err)
	}

	// If the token is empty...
	if token == "" {
		// Check if it was required
		if m.Config.CredentialsOptional {
			// No error, just no token (and that is ok given that CredentialsOptional is true)
			return "OK", nil
		}

		// If we get here, the required token is missing
		errorMsg := "Required authorization token not found"
		return "Unauthorized", fmt.Errorf(errorMsg)
	}

	// Now parse the token

	parsedToken, err := jwt.Parse(token, m.Config.ValidationKeyGetter)
	// Check if there was an error in parsing...
	if err != nil {
		return "Unauthorized", fmt.Errorf("error parsing token: %v", err)
	}

	// If we get here, everything worked.
	ctx.Values().Set(m.Config.ContextKey, parsedToken)
	return "OK", nil
}

// ServeJwt Serve the customized Serve handler for jwt middleware.
func ServeJwt(ctx iris.Context) {
	status, _ := checkJWT(ctx, MyJwtMiddleware)
	ctx.Values().Set("errjwt", status)
	ctx.Next()
}

// GetToken ...
func GetToken(ctx iris.Context) *jwt.Token {
	return MyJwtMiddleware.Get(ctx)
}

// GetUserID returns the user ID parsed from token.
func GetUserID(ctx iris.Context) int64 {
	userToken := GetToken(ctx)
	claims, _ := userToken.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	return int64(id)
}

// CheckLoginStatus is the middleware handler which checks user's login status.
// If "errjwt" in context is "Unauthorized",
// which means some error was raised while parsing JWT,
// that is to say, there is no JWT in current context,
// the handlers of current router will stop immediately.
func CheckLoginStatus(ctx iris.Context) {
	// Error occurs when checking JWT.
	if status := ctx.Values().Get("errjwt"); status == "Unauthorized" {
		response.Unauthorized(ctx, iris.Map{})
		ctx.StopExecution()
		return
	}

	ctx.Next()
}
