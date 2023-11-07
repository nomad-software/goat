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
	"regexp"
	"runtime/cgo"
	"strconv"
	"unsafe"

	"github.com/nomad-software/goat/internal/log"
	"github.com/nomad-software/goat/internal/tk/command"
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
		log.Panic(err, "cannot continue")
	}

	log.Info("initialising the tk package")
	if C.Tk_Init(tk.interpreter) != C.TCL_OK {
		err := tk.getTclError("tk package cannot be initialised")
		log.Panic(err, "cannot continue")
	}

	return tk
}

// Start starts the app main loop. This will immediately show the main window
// and will block until the main window is closed. When this method exits, the
// interpreter is destroyed.
func (tk *Tk) Start() {
	log.Info("starting tk main loop")
	C.Tk_MainLoop()

	log.Info("exited tk main loop")
	tk.Destroy()
}

// Destroy deletes the interpreter and cleans up its resources.
func (tk *Tk) Destroy() {
	log.Info("deleting the interpreter")
	C.Tcl_DeleteInterp(tk.interpreter)
	instance = nil
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
		log.Panic(err, "cannot continue")
	}
}

// GetStrResult gets the interpreter result as a string.
func (tk *Tk) GetStrResult() string {
	result := C.Tcl_GetStringResult(tk.interpreter)
	str := C.GoString(result)

	log.Debug("interpreter result: %v", str)

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

// GetFloatResult gets the interpreter result as a float.
func (tk *Tk) GetFloatResult() float64 {
	str := tk.GetStrResult()

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Error(err)
	}

	return f
}

// GetBoolResult gets the interpreter result as a boolean.
func (tk *Tk) GetBoolResult() bool {
	str := tk.GetStrResult()

	b, err := strconv.ParseBool(str)
	if err != nil {
		log.Error(err)
	}

	return b
}

// SetVarStrValue sets the named variable value using a string.
func (tk *Tk) SetVarStrValue(name string, val string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cval := C.CString(val)
	defer C.free(unsafe.Pointer(cval))

	C.Tcl_SetVar(tk.interpreter, cname, cval, C.TCL_GLOBAL_ONLY)

	log.Debug("set variable {%s} <- {%s}", name, val)
}

// SetVarFloatValue sets the named variable value using a string.
func (tk *Tk) SetVarFloatValue(name string, val float64) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cval := C.CString(fmt.Sprintf("%v", val))
	defer C.free(unsafe.Pointer(cval))

	C.Tcl_SetVar(tk.interpreter, cname, cval, C.TCL_GLOBAL_ONLY)

	log.Debug("set variable {%s} <- {%s}", name, val)
}

// GetVarStrValue gets the named variable value as a string.
func (tk *Tk) GetVarStrValue(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	result := C.Tcl_GetVar(tk.interpreter, cname, C.TCL_GLOBAL_ONLY)
	str := C.GoString(result)

	log.Debug("get variable {%s} -> %s", name, str)

	return str
}

// GetVarIntValue gets the named variable value as an integer.
func (tk *Tk) GetVarIntValue(name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	result := C.Tcl_GetVar(tk.interpreter, cname, C.TCL_GLOBAL_ONLY)
	str := C.GoString(result)

	log.Debug("get variable {%s} -> %s", name, str)

	i, err := strconv.Atoi(str)
	if err != nil {
		log.Error(err)
	}

	return i
}

// GetVarFloatValue gets the named variable value as a float.
func (tk *Tk) GetVarFloatValue(name string) float64 {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	result := C.Tcl_GetVar(tk.interpreter, cname, C.TCL_GLOBAL_ONLY)
	str := C.GoString(result)

	log.Debug("get variable {%s} -> %s", name, str)

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Error(err)
	}

	return f
}

// GetVarBoolValue gets the named variable value as a boolean.
func (tk *Tk) GetVarBoolValue(name string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	result := C.Tcl_GetVar(tk.interpreter, cname, C.TCL_GLOBAL_ONLY)
	str := C.GoString(result)

	log.Debug("get variable {%s} -> %s", name, str)

	b, err := strconv.ParseBool(str)
	if err != nil {
		log.Error(err)
	}

	return b
}

// DeleteVar deletes a variable and cleans up its resources
func (tk *Tk) DeleteVar(name string) {
	log.Debug("deleting variable {%s}", name)

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	result := C.Tcl_UnsetVar(tk.interpreter, cname, C.TCL_GLOBAL_ONLY)
	if result == C.TCL_ERROR {
		err := tk.getTclError("delete variable error: {%s}", name)
		log.Error(err)
	}
}

// CreateCommand creates a custom command in the interpreter.
func (tk *Tk) CreateCommand(name string, callback command.Callback) {
	log.Debug("create command {%s}", name)

	data := &command.CallbackData{
		CommandName: name,
		Callback:    callback,
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	procWrapper := (*[0]byte)(unsafe.Pointer(C.procWrapper))
	delWrapper := (*[0]byte)(unsafe.Pointer(C.delWrapper))
	cdata := C.uintptr_t(cgo.NewHandle(data))

	C.RegisterTclCommand(tk.interpreter, cname, procWrapper, cdata, delWrapper)
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
	data := cgo.Handle(clientData).Value().(*command.CallbackData)

	if argc == 10 {
		data.ElementID = readStringArg(values, 1)
		data.Event.MouseButton = readIntArg(values, 2)
		data.Event.KeyCode = readIntArg(values, 3)
		data.Event.X = readIntArg(values, 4)
		data.Event.Y = readIntArg(values, 5)
		data.Event.Wheel = readIntArg(values, 6)
		data.Event.Key = readStringArg(values, 7)
		data.Event.ScreenX = readIntArg(values, 8)
		data.Event.ScreenY = readIntArg(values, 9)

	} else if argc == 2 {
		data.Dialog.Font = readStringArg(values, 2)
	}

	data.Callback(data)

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
