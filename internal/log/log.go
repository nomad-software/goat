package log

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/nomad-software/goat/internal/thread"
)

func init() {
	if b, ok := os.LookupEnv("SHOW_LOG"); ok {
		if ok, err := strconv.ParseBool(b); err == nil {
			show = ok
		}
	}
}

var (
	show bool
)

// Info logs useful information.
func Info(format string, a ...any) {
	if show {
		format := fmt.Sprintf("[%X] INFO  %s\n", thread.GetTid(), format)
		fmt.Printf(format, a...)
	}
}

// Tcl logs tcl commands when the environment variable is set.
func Tcl(cmd string) {
	if show {
		fmt.Printf("[%X] TCL   %s\n", thread.GetTid(), cmd)
	}
}

// Debug logs useful debug information.
func Debug(format string, a ...any) {
	if show {
		format := fmt.Sprintf("[%X] DEBUG %s\n", thread.GetTid(), format)
		fmt.Printf(format, a...)
	}
}

// Error prints information about the passed error.
func Error(err error) {
	if show {
		fmt.Printf("ERROR %s\n", err)
		for i := 1; i <= 10; i++ {
			_, file, line, _ := runtime.Caller(i)
			if !strings.Contains(file, "goat") {
				break
			}
			fmt.Printf("      - file: %s\n", file)
			fmt.Printf("      - line: %d\n", line)
		}
	}
}

// Panic prints information about the passed error and then panics.
func Panic(err error, msg string) {
	fmt.Printf("PANIC %s\n", err)
	for i := 1; i <= 10; i++ {
		_, file, line, _ := runtime.Caller(i)
		if !strings.Contains(file, "goat") {
			break
		}
		fmt.Printf("      - file: %s\n", file)
		fmt.Printf("      - line: %d\n", line)
	}
	panic(msg)
}
