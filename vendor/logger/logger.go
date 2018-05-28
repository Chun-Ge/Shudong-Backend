package logger

import (
	"os"
	"strings"
	"time"

	"github.com/kataras/iris"

	irisLogger "github.com/kataras/iris/middleware/logger"
)

// Register ..
func Register(app *iris.Application) func() error {
	f := newLogFile()
	close := func() error {
		err := f.Close()
		return err
	}
	app.Logger().SetOutput(f)
	// r, close := newRequestLogger()
	// app.Use(r)
	return close
}

// get a filename based on the date, file logs works that way the most times
// but these are just a sugar.
func todayFilename() string {
	today := time.Now().Format("Jan-02-2006")
	return today + ".log"
}

func newLogFile() *os.File {
	filename := todayFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}

func newRequestLogger() (h iris.Handler, close func() error) {
	close = func() error { return nil }

	c := irisLogger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}

	logFile := newLogFile()
	close = func() error {
		err := logFile.Close()
		return err
	}

	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := irisLogger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message, headerMessage)
		logFile.Write([]byte(output))
	}

	// we don't want to use the logger
	// to log requests to assets and etc
	c.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})

	h = irisLogger.New(c)

	return
}
