package tk

/*
#cgo CFLAGS: -I/usr/include/tcl/
#cgo LDFLAGS: -ltcl
#cgo LDFLAGS: -ltk

#include <stdlib.h>
#include <stdint.h>

#if __has_include(<tcl/tk.h>)
  #include <tcl/tk.h>

#elif __has_include(<tk.h>)
  #include <tk.h>

#else
  #error "tk.h not found"

#endif

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
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"runtime/cgo"
	"strconv"
	"strings"
	"sync"
	"unsafe"

	"github.com/nomad-software/goat/command"
	"github.com/nomad-software/goat/internal/log"
	"github.com/nomad-software/goat/internal/thread"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

var (
	// Global interpreter instance.
	instance *Tk
	once     sync.Once

	Binding      = regexp.MustCompile(`^<.*?>$`)
	Event        = regexp.MustCompile(`^<.*?>$`)
	VirtualEvent = regexp.MustCompile(`^<<.*?>>$`)
)

// Get gets the global instance of the interpreter.
func Get() *Tk {
	once.Do(func() {
		instance = new()
	})

	return instance
}

// Tk is the main interpreter.
type Tk struct {
	interpreter *C.Tcl_Interp // The low level C based interpreter.
	tid         uint64        // Thread id of the interpreter.
	queue       chan func()   // Channel to send command on if they are on a different thread.
}

// new creates a new instance of the interpreter.
// This will end the program on any error.
func new() *Tk {
	runtime.LockOSThread() // Lock Tcl/Tk calls to one thread.

	log.Info("creating new interpreter")

	tk := &Tk{
		interpreter: C.Tcl_CreateInterp(),
		tid:         thread.GetTid(),
		// A buffered channel is crucial for this interpreter design. If a
		// command is sent to the interpreter from a different thread, we add
		// that command to the queue and it is marshalled to the correct thread
		// and then executed. It's not a great design because the buffer size
		// is a guess of how many commands we would need to support from other
		// threads. I can refine this later.
		queue: make(chan func(), 4096),
	}

	log.Info("interpreter locked to thread: %X", tk.tid)

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

	if tk.tid != thread.GetTid() {
		log.Panic(errors.New("start must be called on the interpreter thread"), "cannot continue")
	}

	for C.Tk_GetNumMainWindows() > 0 {
		select {
		case fn := <-tk.queue:
			if fn != nil {
				fn()
			}
		default:
			C.Tcl_DoOneEvent(0)
		}
	}

	log.Info("exited main loop")

	tk.Destroy()
}

// Destroy deletes the interpreter and cleans up its resources.
func (tk *Tk) Destroy() {
	if tk == nil {
		return
	}

	tk.run(func() {
		log.Info("deleting the interpreter")

		if tk.interpreter != nil {
			C.Tcl_DeleteInterp(tk.interpreter)
			tk.interpreter = nil
		}

		close(tk.queue)
		for range tk.queue {
		}

		instance = nil
		runtime.UnlockOSThread() // Unlock the Tcl/Tk thread.
	})
}

// Eval passes the specified command to the interpreter for evaluation.
// This will end the program on any error.
func (tk *Tk) Eval(format string, a ...any) {
	if tk == nil {
		return
	}

	tk.run(func() {
		cmd := fmt.Sprintf(format, a...)
		cstr := C.CString(cmd)
		defer C.free(unsafe.Pointer(cstr))

		log.Tcl(cmd)

		result := C.Tcl_EvalEx(tk.interpreter, cstr, -1, 0)

		if result == C.TCL_ERROR {
			err := tk.getTclError("evaluation error")
			log.Error(err)
		}
	})
}

// run makes sure functions passed run in the interpreter's thread.
func (tk *Tk) run(fn func()) {
	if tk == nil {
		return
	}

	if tk.tid == thread.GetTid() {
		fn()
		return
	}

	tk.queue <- fn
}

// GetStrResult gets the interpreter result as a string.
func (tk *Tk) GetStrResult() string {
	var str string

	tk.run(func() {
		result := C.Tcl_GetStringResult(tk.interpreter)
		str = C.GoString(result)

		log.Debug("interpreter result: %v", str)
	})

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

func (tk *Tk) GetStrSliceResult() []string {
	str := tk.GetStrResult()

	return parseTclList(str)
}

// SetVarStrValue sets the named variable value using a string.
func (tk *Tk) SetVarStrValue(name string, val string) {
	tk.run(func() {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		cval := C.CString(val)
		defer C.free(unsafe.Pointer(cval))

		C.Tcl_SetVar(tk.interpreter, cname, cval, C.TCL_GLOBAL_ONLY)

		log.Debug("set variable {%s} <- {%s}", name, val)
	})
}

// SetVarFloatValue sets the named variable value using a string.
func (tk *Tk) SetVarFloatValue(name string, val float64) {
	tk.run(func() {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		cval := C.CString(fmt.Sprintf("%v", val))
		defer C.free(unsafe.Pointer(cval))

		C.Tcl_SetVar(tk.interpreter, cname, cval, C.TCL_GLOBAL_ONLY)

		log.Debug("set variable {%s} <- {%v}", name, val)
	})
}

// GetVarStrValue gets the named variable value as a string.
func (tk *Tk) GetVarStrValue(name string) string {
	var str string

	tk.run(func() {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		result := C.Tcl_GetVar(tk.interpreter, cname, C.TCL_GLOBAL_ONLY)
		str = C.GoString(result)

		log.Debug("get variable {%s} -> %s", name, str)
	})

	return str
}

// GetVarIntValue gets the named variable value as an integer.
func (tk *Tk) GetVarIntValue(name string) int {
	str := tk.GetVarStrValue(name)

	i, err := strconv.Atoi(str)
	if err != nil {
		log.Error(err)
	}

	return i
}

// GetVarFloatValue gets the named variable value as a float.
func (tk *Tk) GetVarFloatValue(name string) float64 {
	str := tk.GetVarStrValue(name)

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Error(err)
	}

	return f
}

// GetVarBoolValue gets the named variable value as a boolean.
func (tk *Tk) GetVarBoolValue(name string) bool {
	str := tk.GetVarStrValue(name)

	b, err := strconv.ParseBool(str)
	if err != nil {
		log.Error(err)
	}

	return b
}

// DestroyVar destroys a variable and cleans up its resources.
func (tk *Tk) DestroyVar(name string) {
	tk.run(func() {
		log.Debug("deleting variable {%s}", name)

		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		result := C.Tcl_UnsetVar(tk.interpreter, cname, C.TCL_GLOBAL_ONLY)
		if result == C.TCL_ERROR {
			err := tk.getTclError("destroy variable error: {%s}", name)
			log.Error(err)
		}
	})
}

// CreateCommand creates a custom command in the interpreter.
func (tk *Tk) CreateCommand(el element.Element, name string, callback command.Callback) {
	tk.run(func() {
		log.Debug("create command {%s}", name)

		data := &command.CommandData{
			Element:     el,
			CommandName: name,
			Callback:    callback,
		}

		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		procWrapper := (*[0]byte)(unsafe.Pointer(C.procWrapper))
		delWrapper := (*[0]byte)(unsafe.Pointer(C.delWrapper))
		cdata := C.uintptr_t(cgo.NewHandle(data))

		C.RegisterTclCommand(tk.interpreter, cname, procWrapper, cdata, delWrapper)
	})
}

// CreateBindCommand creates a custom command in the interpreter.
func (tk *Tk) CreateBindCommand(el element.Element, name string, callback command.BindCallback) {
	tk.run(func() {
		log.Debug("create bind command {%s}", name)

		data := &command.BindData{
			Element:     el,
			CommandName: name,
			Callback:    callback,
		}

		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		procWrapper := (*[0]byte)(unsafe.Pointer(C.procWrapper))
		delWrapper := (*[0]byte)(unsafe.Pointer(C.delWrapper))
		cdata := C.uintptr_t(cgo.NewHandle(data))

		C.RegisterTclCommand(tk.interpreter, cname, procWrapper, cdata, delWrapper)
	})
}

// CreateFontDialogCommand creates a custom command in the interpreter.
func (tk *Tk) CreateFontDialogCommand(el element.Element, name string, callback command.FontDialogCallback) {
	tk.run(func() {
		log.Debug("create font dialog command {%s}", name)

		data := &command.FontData{
			Element:     el,
			CommandName: name,
			Callback:    callback,
		}

		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		procWrapper := (*[0]byte)(unsafe.Pointer(C.procWrapper))
		delWrapper := (*[0]byte)(unsafe.Pointer(C.delWrapper))
		cdata := C.uintptr_t(cgo.NewHandle(data))

		C.RegisterTclCommand(tk.interpreter, cname, procWrapper, cdata, delWrapper)
	})
}

// DestroyCommand destroys a command and cleans up its resources.
func (tk *Tk) DestroyCommand(name string) {
	tk.run(func() {
		log.Debug("destroy command {%s}", name)

		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		status := C.Tcl_DeleteCommand(tk.interpreter, cname)
		if status != C.TCL_OK {
			err := tk.getTclError("destroy command failed")
			log.Error(err)
		}
	})
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

	switch data := cgo.Handle(clientData).Value().(type) {
	case *command.CommandData:
		log.Debug("command data: %#v", data)
		data.Callback(data)

	case *command.BindData:
		data.Event.MouseButton = readIntArg(values, 1)
		data.Event.KeyCode = readIntArg(values, 2)
		data.Event.ElementX = readIntArg(values, 3)
		data.Event.ElementY = readIntArg(values, 4)
		data.Event.Wheel = readIntArg(values, 5)
		data.Event.Key = readStringArg(values, 6)
		data.Event.ScreenX = readIntArg(values, 7)
		data.Event.ScreenY = readIntArg(values, 8)
		log.Debug("bind data: %#v", data)
		data.Callback(data)

	case *command.FontData:
		data.Font = parseFontDialogResult(readStringArg(values, 1))

		log.Debug("font data: %#v", data)
		data.Callback(data)
	}

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

// parseTclList parses a tcl list from the interpreter result and tried to
// correctly handle any curly brackets that may be emmited for escaping
// purposes.
func parseTclList(str string) []string {
	result := make([]string, 0)

	var out string
	var count int

	for _, r := range str {
		if r == '{' {
			count += 1
		}
		if r == '}' {
			count -= 1
		}
		if r == ' ' && count == 0 {
			result = append(result, out)
			out = ""
			continue
		}
		out += string(r)
	}

	if out != "" {
		result = append(result, out)
	}

	for i, str := range result {
		if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {
			result[i] = str[1 : len(str)-1]
		}
	}

	return result
}

func parseFontDialogResult(str string) command.Font {
	font := command.Font{}

	details := parseTclList(str)

	if len(details) > 0 {
		font.Name = details[0]
	}

	if len(details) > 1 {
		font.Size = details[1]
	}

	if len(details) > 2 {
		font.Modifiers = details[2:]
	}

	return font
}
