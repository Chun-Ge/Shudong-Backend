package logger

import (
	"fmt"
	"net/http"
	"os"

	"args"
	"err"
	"utils"

	"github.com/kataras/iris"
	"github.com/robfig/cron"
)

// CurrentLogFile ...
var CurrentLogFile *os.File

// Register ...
func Register(app *iris.Application) {
	generalErrLog(app)

	updateLogFilenameDaily(app)
}

func generalErrLog(app *iris.Application) {
	app.OnAnyErrorCode(func(ctx iris.Context) {
		if args.DEBUG {
			ctx.Writef("request route: %+v\n", ctx.Path())
		}
		statusCode := ctx.GetStatusCode()
		ctx.Writef("%+v %+v\n", statusCode, http.StatusText(statusCode))
		app.Logger().Warnf(ctx.Path())
	})
}

func updateLogFilenameDaily(app *iris.Application) {
	CurrentLogFile = newLogFile()
	app.Logger().SetOutput(CurrentLogFile)

	c := cron.New()
	c.AddFunc("@midnight", func() {
		er := utils.GetCloseFileFunc(CurrentLogFile)()
		err.CheckErrWithPanic(er)

		CurrentLogFile = newLogFile()
		if args.DEBUG {
			fmt.Printf("Change Log File to %+v\n", CurrentLogFile.Name())
		}
		// app.Logger().Infof("Change Log File to %+v\n", CurrentLogFile.Name())
		app.Logger().SetOutput(CurrentLogFile)
	})
	c.Start()
}
