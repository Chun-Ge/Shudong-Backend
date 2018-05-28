package server

import (
	"args"
	stdContext "context"
	"logger"
	"middlewares"
	"os"
	"os/signal"
	"route"
	"syscall"
	"time"

	"github.com/kataras/iris"
)

// Start .
func Start() {
	app := iris.New()

	middlewares.Register(app)
	route.Register(app)

	app.Run(iris.Addr("" + args.Port))
}

// StartWithConfiguration reads the config file and
func StartWithConfiguration(configFilePath string) {
	app := iris.New()

	middlewares.Register(app)
	route.Register(app)

	close := logger.Register(app)
	defer close()

	app.Configure(iris.WithConfiguration(iris.YAML(configFilePath)))

	go withGracefulShutdown(app)()

	app.Run(iris.Addr("" + args.Port))
}

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
