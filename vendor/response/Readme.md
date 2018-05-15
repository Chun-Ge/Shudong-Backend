# Package response

This module provides much simpler methods for responses by wrapping `ctx.StatusCode()` and `ctx.JSON()`.

Example:

```go
import "response"
// ...
func exampleHandler(ctx iris.Context) {
    // retData can be any data type (interface{})
    var testID int64 = 3
    retData := model.GetPostByID(testID) // this time it will be *entity.Post
    response.OK(ctx, retData)
}
```