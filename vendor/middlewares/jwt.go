package middlewares

import (
	"args"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

// checkJWT the main functionality, checks for token.
func checkJWT(ctx iris.Context, m *jwtmiddleware.Middleware) error {
	if !m.Config.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return nil
		}
	}

	token, err := m.Config.Extractor(ctx)

	if err != nil {
		return fmt.Errorf("Error extracting token: %v", err)
	}

	// If the token is empty...
	if token == "" {
		// Check if it was required
		if m.Config.CredentialsOptional {
			// No error, just no token (and that is ok given that CredentialsOptional is true)
			return nil
		}

		// If we get here, the required token is missing
		errorMsg := "Required authorization token not found"
		return fmt.Errorf(errorMsg)
	}

	// Now parse the token

	parsedToken, err := jwt.Parse(token, m.Config.ValidationKeyGetter)
	// Check if there was an error in parsing...
	if err != nil {
		return fmt.Errorf("error parsing token: %v", err)
	}

	if m.Config.SigningMethod != nil && m.Config.SigningMethod.Alg() != parsedToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			m.Config.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		return fmt.Errorf("Error validating token algorithm: %s", message)
	}

	// Check if the parsed token is valid...
	if !parsedToken.Valid {
		return fmt.Errorf("Token is invalid")
	}

	if m.Config.Expiration {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			if expired := claims.VerifyExpiresAt(time.Now().Unix(), true); !expired {
				return fmt.Errorf("Token is expired")
			}
		}
	}

	// If we get here, everything worked.
	ctx.Values().Set(m.Config.ContextKey, parsedToken)

	return nil
}

// Serve the customized Serve handler for jwt middleware.
func ServeJWT(ctx iris.Context) {
	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(args.SecretKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ContextKey:    "jwt",
		Expiration:    true,
	})

	if err := checkJWT(ctx, myJwtMiddleware); err != nil {
		ctx.Values().Set("errjwt", err)
	} else {
		ctx.Values().Set("errjwt", "")
	}
	ctx.Next()
}
