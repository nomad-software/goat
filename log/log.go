package log

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

// Info logs useful information.
func Info(format string, a ...any) {
	fmt.Printf("INFO  "+format+"\n", a...)
}

// Tcl logs tcl commands when the environment variable is set.
func Tcl(cmd string) {
	if b, ok := os.LookupEnv("LOG_TCL"); ok {
		if ok, err := strconv.ParseBool(b); err == nil {
			if ok {
				fmt.Printf("TCL   %s\n", cmd)
			}
		}
	}
}

// Debug logs useful debug information.
func Debug(format string, a ...any) {
	fmt.Printf("DEBUG "+format+"\n", a...)
}

// Error prints information about the passed error.
func Error(err error) {
	fmt.Printf("ERROR %s\n", err)

	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("      - file: %s\n", file)
	fmt.Printf("      - line: %d\n", line)

	_, file, line, _ = runtime.Caller(2)
	fmt.Printf("      - caller: %s\n", file)
	fmt.Printf("      - line: %d\n", line)
}
