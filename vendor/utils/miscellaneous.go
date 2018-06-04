package utils

import (
	"fmt"

	"args"
	"os"
)

// GetCloseFileFunc ...
func GetCloseFileFunc(f *os.File) func() error {
	return func() error {
		if args.DEBUG {
			fmt.Printf("Close %+v\n", f.Name())
		}
		err := f.Close()
		return err
	}
}
