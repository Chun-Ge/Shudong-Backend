package logger

import (
	"os"
	"path/filepath"

	"args"
	"err"
	"utils"
)

// get a filename based on the date, file logs works that way the most times
// but these are just a sugar.
func todayFilename() string {
	// today := time.Now().Format("2018-May-25")
	today := utils.GetDateToday()
	return today + ".log"
}

func newLogFile() *os.File {
	filename := filepath.Join(args.LogDir(), todayFilename())
	// open an output file, when it comes to another day, a cron job will
	// call this func again and attach app.Logger() to a new file
	f, er := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	err.CheckErrWithPanic(er)

	return f
}
