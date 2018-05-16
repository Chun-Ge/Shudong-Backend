package service

import (
	"args"
	"model"
	"response"

	"github.com/kataras/iris"

	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

type postInfo struct {
	UserID     int64
	CategoryID int64  `form:"category"`
	Title      string `form:"title"`
	Content    string `form:"content"`
}

// CreatePost Post a post.
func CreatePost(ctx iris.Context) {
	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(args.SecretKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ContextKey:    "jwt",
		Expiration:    true,
	})

	userToken := myJwtMiddleware.Get(ctx)
	if claims, ok := userToken.Claims.(jwt.MapClaims); ok && userToken.Valid {
		// claimTestedValue = claims["foo"].(string)
		userID := claims["id"].(int64)

		info := postInfo{UserID: userID}
		ctx.ReadForm(&info)

		post, err := model.NewPostWithRandomName(info.UserID, info.CategoryID, info.Title, info.Content)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError) // code: 500
			ctx.JSON(iris.Map{
				"msg": "InternalServerError",
				//"data" : ""
			})
			return
		}

		response.OK(ctx, iris.Map{
			"msg": "OK",
			"data": iris.Map{
				"post": iris.Map{
					"postId":       post.ID,
					"authorId":     post.UserID,
					"title":        post.Title,
					"content":      post.Content,
					"likeCount":    post.Like,
					"commentCount": 0, //
				},
			},
		})
	} else {
		response.InternalServerError(ctx, iris.Map{
			"msg": "InternalServerError",
			//"data" : ""
		})
	}
}
