package utils

import (
	"fmt"
	"os"
)

func ErrorRecorder(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		// os.Exit(1)
	}
}
