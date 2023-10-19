package tk

/*
#cgo CFLAGS: -I/usr/include/tcl/
#cgo LDFLAGS: -ltcl
#cgo LDFLAGS: -ltk

#include <stdlib.h>
#include <tcl/tk.h>
*/
import "C"

import (
	"fmt"
	"log/slog"
	"os"
	"unsafe"

	_ "github.com/nomad-software/goat/log"
)

// Tcl_CreateInterp
// Tcl_Init
// Tcl_DeleteInterp

// Tk_Init
// Tk_MainLoop

// Tcl_EvalEx
// Tcl_GetStringResult
// Tcl_SetResult

// Tcl_SetVar
// Tcl_GetVar
// Tcl_UnsetVar
// Tcl_CreateCommand
// Tcl_DeleteCommand

var instance *Tk // Global interpreter instance.

// Get gets the global instance of the interpreter.
func Get() *Tk {
	if instance != nil {
		return instance
	}

	instance = new()

	return instance
}

// Tk is the main interpreter.
type Tk struct {
	interpreter *C.Tcl_Interp // The low level C based interpreter.
}

// new creates a new instance of the interpreter.
// This will end the program on any error.
func new() *Tk {
	slog.Info("creating new interpreter")

	tk := &Tk{
		interpreter: C.Tcl_CreateInterp(),
	}

	slog.Info("initialising interpreter")
	if C.Tcl_Init(tk.interpreter) != C.TCL_OK {
		err := tk.getTclError("interpreter cannot be initialised")
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("initialising the tk package")
	if C.Tk_Init(tk.interpreter) != C.TCL_OK {
		err := tk.getTclError("tk package cannot be initialised")
		slog.Error(err.Error())
		os.Exit(1)
	}

	return tk
}

// Start starts the tk main loop.
// This will immediately show the room window.
func (tk *Tk) Start() {
	slog.Info("starting the tk main loop")
	C.Tk_MainLoop() // This will block until the main window is closed.

	slog.Info("freeing the interpreter")
	C.Tcl_DeleteInterp(tk.interpreter)
}

// Eval passes the specified command to the interpreter for evaluation.
// This will end the program on any error.
func (tk *Tk) Eval(format string, a ...any) {
	cmd := fmt.Sprintf(format, a...)

	slog.Debug("tcl", "", cmd)

	cstr := C.CString(cmd)
	defer C.free(unsafe.Pointer(cstr))

	result := C.Tcl_EvalEx(tk.interpreter, cstr, -1, 0)

	if result == C.TCL_ERROR {
		err := tk.getTclError("evaluation error")
		slog.Error(err.Error())
		os.Exit(1)
	}
}

// createError reads the last result from the interpreter and returns it as
// a normal Go error.
func (tk *Tk) getTclError(format string, a ...any) error {
	result := C.Tcl_GetStringResult(tk.interpreter)
	str := C.GoString(result)
	err := fmt.Errorf("%s: %s", fmt.Sprintf(format, a...), str)
	return err
}
