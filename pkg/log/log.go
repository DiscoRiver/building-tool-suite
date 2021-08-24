package log

import (
	"fmt"
	"os"
)

var (
	Verbose bool
)

// Info logs a message as INFO
func Info(format string, args ...interface{}) {
	fmt.Printf("INFO: "+format, args...)
}

// Debug logs a message as DEBUG, only if Verbose is true.
func Debug(format string, args ...interface{}) {
	if Verbose {
		fmt.Printf("DEBUG: "+format, args...)
	}
}

// Error logs a message as ERROR, then exits the program with exit code 1.
func Error(format string, args ...interface{}) {
	fmt.Printf("ERROR: "+format, args...)
	os.Exit(1)
}
