package log

import "fmt"

var (
	Verbose bool
)

func Info(format string, args ...interface{}) {
	fmt.Printf("INFO: "+format, args...)
}

func Debug(format string, args ...interface{}) {
	if Verbose {
		fmt.Printf("DEBUG: "+format, args...)
	}
}

func Error(format string, args ...interface{}) {
	fmt.Printf("ERROR: "+format, args...)
}
