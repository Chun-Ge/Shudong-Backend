package err

import (
	"fmt"
)

// Callback .
type Callback func()

const (
	// SQLUpdateError .
	SQLUpdateError = "SQL Update Error"
)

// CheckErr checks if an error exists and calls panic() if so.
//
// Deprecated: Use CheckErrWithPanic() instead.
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

// CheckErrWithPanic is simply an alias of CheckErr(), to avoid unnecessary changes.
// Use CheckErrWithPanic if panic() is needed.
func CheckErrWithPanic(e error) {
	CheckErr(e)
}

// CheckErrWithCallback checks if an error exists
// and calls callback() if so.
func CheckErrWithCallback(e error, callback Callback) {
	if e != nil {
		callback()
	}
}

// CheckErrWithPrintln ..
func CheckErrWithPrintln(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}
