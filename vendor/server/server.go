package server

import (
	stdContext "context"
	"database"
	"fmt"
	"os"
	"os/signal"
	"response"
	"syscall"
	"time"

	"args"
	"logger"
	"middlewares"
	"route"

	"github.com/kataras/iris"
)

// Start ...
func Start() {
	// start database
	database.Start()

	app := iris.New()

	middlewares.Register(app)
	route.Register(app)

	app.Run(iris.Addr("" + args.Port))
}

// StartWithConfiguration starts the app according to the config file.
func StartWithConfiguration(configFilePath string) {
	// start database
	database.Start()

	// app := iris.New()
	app := iris.Default()

	// register all middlewares
	middlewares.Register(app)

	// register all routes
	route.Register(app)

	// setup app.Logger()
	logger.Register(app)

	// the last CurrentLogFile cannot be closed by
	// defer utils.GetCloseFileFunc(logger.CurrentLogFile)()
	// because the closure generated at that time will store
	// the very first filename of the log.
	// (e.g. Day1.log but Day7 is expected)
	defer func() {
		if args.DEBUG {
			fmt.Printf("Close %+v\n", logger.CurrentLogFile.Name())
		}
		logger.CurrentLogFile.Close()
	}()

	app.OnErrorCode(500, func(ctx iris.Context) {
		response.InternalServerError(ctx, iris.Map{})
	})

	// Configurations
	app.Configure(iris.WithConfiguration(iris.YAML(configFilePath)))

	go withGracefulShutdown(app)()

	writeDebugModeToLog()

	app.Run(iris.Addr(":" + args.Port))
}

// Graceful Shutdown: use a goroutine
// to catch os.Interrupt, os.Kill, SIGINT, SIGKILL, SIGTERM
// and then call app.Shutdown()
func withGracefulShutdown(app *iris.Application) func() {
	return func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			// kill -SIGINT XXXX or Ctrl+c
			os.Interrupt,
			syscall.SIGINT, // register that too, it should be ok
			// os.Kill  is equivalent with the syscall.Kill
			os.Kill,
			syscall.SIGKILL, // register that too, it should be ok
			// kill -SIGTERM XXXX
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			println("shutdown...")

			timeout := 5 * time.Second
			ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
			defer cancel()
			app.Shutdown(ctx)
		}
	}
}

func writeDebugModeToLog() {
	logString := "\n[START] Shudong-Backend in "
	if args.DEBUG {
		logString += "[Debug Mode]...\n"
	} else {
		logString += "[Production Mode]...\n"
	}
	logger.CurrentLogFile.WriteString(logString)
}
