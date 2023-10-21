package tk

/*
#cgo CFLAGS: -I/usr/include/tcl/
#cgo LDFLAGS: -ltcl
#cgo LDFLAGS: -ltk

#include <stdlib.h>
#include <stdint.h>
#include <tcl/tk.h>

#if _WIN32
	int __declspec(dllexport) procWrapper(ClientData clientData, Tcl_Interp* interp, int argc, char** argv);
	void __declspec(dllexport) delWrapper(ClientData clientData);
#else
	int procWrapper(ClientData clientData, Tcl_Interp* interp, int argc, char** argv);
	void delWrapper(ClientData clientData);
#endif

static void RegisterTclCommand(Tcl_Interp* interp, char* name, int (*proc)(ClientData, Tcl_Interp*, int, const char**), uintptr_t clientData, void (*del)(ClientData)) {
    Tcl_CreateCommand(interp, name, proc, (ClientData)clientData, del);
}

*/
import "C"

import (
	"fmt"
	"os"
	"regexp"
	"runtime/cgo"
	"strconv"
	"unsafe"

	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/tk/command"
)

var (
	// Global interpreter instance.
	instance *Tk

	Binding      = regexp.MustCompile(`^<.*?>$`)
	Event        = regexp.MustCompile(`^<.*?>$`)
	VirtualEvent = regexp.MustCompile(`^<<.*?>>$`)
)

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

// CreateCommand creates a custom command in the interpreter.
func (tk *Tk) CreateCommand(name string, callback command.Callback) {
	log.Debug("create command {%s}", name)

	payload := &command.CallbackPayload{
		CommandName: name,
		Callback:    callback,
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	procWrapper := (*[0]byte)(unsafe.Pointer(C.procWrapper))
	delWrapper := (*[0]byte)(unsafe.Pointer(C.delWrapper))
	cpayload := C.uintptr_t(cgo.NewHandle(payload))

	C.RegisterTclCommand(tk.interpreter, cname, procWrapper, cpayload, delWrapper)
}

// DeleteCommand deletes the specified command from the interpreter.
func (tk *Tk) DeleteCommand(name string) {
	log.Debug("delete command {%s}", name)

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	status := C.Tcl_DeleteCommand(tk.interpreter, cname)
	if status != C.TCL_OK {
		err := tk.getTclError("delete command failed")
		log.Error(err)
	}
}

// getTclError reads the last result from the interpreter and returns it as
// a normal Go error.
func (tk *Tk) getTclError(format string, a ...any) error {
	str := tk.GetStrResult()
	err := fmt.Errorf("%s: %s", fmt.Sprintf(format, a...), str)
	return err
}

// procWrapper is an exported C ABI function to make interop a little easier.
// This function is called when a bound event fires.
//
//export procWrapper
func procWrapper(clientData unsafe.Pointer, interp *C.Tcl_Interp, argc C.int, argv **C.char) C.int {
	values := unsafe.Slice(argv, argc)
	payload := cgo.Handle(clientData).Value().(*command.CallbackPayload)

	if argc == 10 {
		payload.ElementID = readStringArg(values, 1)
		payload.Event.MouseButton = readIntArg(values, 2)
		payload.Event.KeyCode = readIntArg(values, 3)
		payload.Event.X = readIntArg(values, 4)
		payload.Event.Y = readIntArg(values, 5)
		payload.Event.Wheel = readIntArg(values, 6)
		payload.Event.Key = readStringArg(values, 7)
		payload.Event.ScreenX = readIntArg(values, 8)
		payload.Event.ScreenY = readIntArg(values, 9)

	} else if argc == 2 {
		payload.Dialog.Font = readStringArg(values, 2)
	}

	payload.Callback(payload)

	return C.TCL_OK
}

// delWrapper is an exported C ABI function to make interop a little easier.
// This function is called when a command is deleted.
//
//export delWrapper
func delWrapper(clientData unsafe.Pointer) {
	cgo.Handle(clientData).Delete()
}

// readIntArg is a helper function to read int based arguments passed to the
// procWrapper.
func readIntArg(argv []*C.char, index int) int {
	val := C.GoString(argv[index])
	if val == "??" {
		return 0
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		log.Error(err)
		return 0
	}

	return i
}

// readIntArg is a helper function to read string based arguments passed to the
// procWrapper.
func readStringArg(argv []*C.char, index int) string {
	val := C.GoString(argv[index])
	if val == "??" {
		return ""
	}
	return val
}
