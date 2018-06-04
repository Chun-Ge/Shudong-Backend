package utils

import (
	"args"

	"time"
)

var (
	dateToday string
)

func init() {
	setDateToday(getDateNow())
}

// GetDateToday ...
func GetDateToday() string {
	checkDateAndUpdate()
	return dateToday
}

func checkDateAndUpdate() {
	dateNow := getDateNow()
	if dateNow != dateToday {
		setDateToday(dateNow)
	}
}

func getDateNow() string {
	return time.Now().Format(args.TimeFormat)
}

func setDateToday(newDate string) {
	dateToday = newDate
}
