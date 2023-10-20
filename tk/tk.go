package tk

/*
#cgo CFLAGS: -I/usr/include/tcl/
#cgo LDFLAGS: -ltcl
#cgo LDFLAGS: -ltk

#include <stdlib.h>
#include <tcl/tk.h>

int TclCustomCommand(void* clientData, Tcl_Interp* interp, int argc, const char* argv) {

    return TCL_OK;
}
*/
import "C"

import (
	"fmt"
	"os"
	"strconv"
	"unsafe"

	"github.com/nomad-software/goat/log"
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
	log.Info("creating new interpreter")

	tk := &Tk{
		interpreter: C.Tcl_CreateInterp(),
	}

	log.Info("initialising interpreter")
	if C.Tcl_Init(tk.interpreter) != C.TCL_OK {
		err := tk.getTclError("interpreter cannot be initialised")
		log.Error(err)
		tk.Destroy(1)
	}

	log.Info("initialising the tk package")
	if C.Tk_Init(tk.interpreter) != C.TCL_OK {
		err := tk.getTclError("tk package cannot be initialised")
		log.Error(err)
		tk.Destroy(1)
	}

	return tk
}

// Start starts the tk main loop.
// This will immediately show the room window.
func (tk *Tk) Start() {
	log.Info("starting tk main loop")
	C.Tk_MainLoop() // This will block until the main window is closed.

	log.Info("exited tk main loop")
	tk.Destroy(0)
}

// Destroy deletes the interpreter and cleans up its resources.
func (tk *Tk) Destroy(code int) {
	log.Info("deleting the interpreter")
	C.Tcl_DeleteInterp(tk.interpreter)

	os.Exit(code)
}

// Eval passes the specified command to the interpreter for evaluation.
// This will end the program on any error.
func (tk *Tk) Eval(format string, a ...any) {
	cmd := fmt.Sprintf(format, a...)

	log.Tcl(cmd)

	cstr := C.CString(cmd)
	defer C.free(unsafe.Pointer(cstr))

	result := C.Tcl_EvalEx(tk.interpreter, cstr, -1, 0)

	if result == C.TCL_ERROR {
		err := tk.getTclError("evaluation error")
		log.Error(err)
		tk.Destroy(1)
	}
}

// GetStrResult gets the interpreter result as a string.
func (tk *Tk) GetStrResult() string {
	result := C.Tcl_GetStringResult(tk.interpreter)
	str := C.GoString(result)

	return str
}

// GetStrResult gets the interpreter result as a string.
func (tk *Tk) GetIntResult() int {
	str := tk.GetStrResult()

	i, err := strconv.Atoi(str)
	if err != nil {
		log.Error(err)
	}

	return i
}

func (tk *Tk) CreateCommand(name string) {
	cname := C.CString(name)

	C.Tcl_CreateCommand(tk.interpreter, cname, C.TclCustomCommand, nil, nil)
}

// createError reads the last result from the interpreter and returns it as
// a normal Go error.
func (tk *Tk) getTclError(format string, a ...any) error {
	str := tk.GetStrResult()
	err := fmt.Errorf("%s: %s", fmt.Sprintf(format, a...), str)
	return err
}
