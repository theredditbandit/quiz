// this package will just contain utility functions that will be (re)used all over the program

package utils

import (
	"fmt"
	"os"
)

// print an exit message and exit with code 1
func ExitWithMessage(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}
