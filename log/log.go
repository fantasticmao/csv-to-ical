package log

import (
	"fmt"
	"os"
)

func Panic(err error, format string, args ...interface{}) {
	Error(format, args...)
	panic(err)
}

func Info(format string, args ...interface{}) {
	_, _ = fmt.Printf(format, args...)
	_, _ = fmt.Println()
}

func Error(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	_, _ = fmt.Fprintln(os.Stderr)
}
