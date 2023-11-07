package utils

import (
	"fmt"
	"os"
)

// print an exit message and exit with code 1
func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
