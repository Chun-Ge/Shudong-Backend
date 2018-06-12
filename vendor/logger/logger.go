package logger

import (
	"fmt"
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
	updateLogFilenameDaily(app)
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
