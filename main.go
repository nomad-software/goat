// package main
//
// import (
// 	"github.com/nomad-software/goat/app"
// 	"github.com/nomad-software/goat/log"
// 	"github.com/nomad-software/goat/window"
// )
//
// func main() {
// 	app := app.New()
// 	main := app.GetMainWindow()
//
// 	main.SetSize(1024, 768)
// 	main.SetTitle("test")
//
// 	win2 := window.New(main)
// 	win2.SetTitle("test2")
// 	win2.Raise(main)
// 	win2.Focus(true)
//
// 	log.Debug("main win cursor y pos: %v", main.GetCursorYPos())
//
// 	log.Debug("screen width: %v", main.GetScreenWidth())
// 	log.Debug("screen height: %v", main.GetScreenHeight())
//
// 	log.Debug("x pos: %v", main.GetXPos(false))
// 	log.Debug("x pos: %v", main.GetXPos(true))
//
// 	log.Debug("y pos: %v", main.GetYPos(false))
// 	log.Debug("y pos: %v", main.GetYPos(true))
//
// 	app.Start()
// }

package main

/*
#cgo CFLAGS: -I/usr/include/tcl/
#cgo LDFLAGS: -ltcl
#cgo LDFLAGS: -ltk

#include <stdlib.h>
#include <stdint.h>
#include <tcl/tk.h>

int CommandWrapper(ClientData clientData, Tcl_Interp* interp, int argc, char** argv);

#if _WIN32
	int __declspec(dllexport) CommandWrapper(ClientData clientData, Tcl_Interp* interp, int argc, char** argv);
#else
	int CommandWrapper(ClientData clientData, Tcl_Interp* interp, int argc, char** argv);
#endif

static int RegisterTclCommand(Tcl_Interp* interp, char* name, int (*func)(ClientData, Tcl_Interp*, int, const char**), uintptr_t clientData) {
    Tcl_CreateCommand(interp, name, func, (ClientData)clientData, NULL);
	return TCL_OK;
}*/
import "C"

import (
	"fmt"
	"runtime/cgo"
	"unsafe"

	"github.com/nomad-software/goat/tk/command"
)

//export CommandWrapper
func CommandWrapper(clientData unsafe.Pointer, interp *C.Tcl_Interp, argc C.int, argv **C.char) C.int {
	fmt.Printf("Number of arguments: %d\n", argc)
	fmt.Println("Arguments:")
	values := unsafe.Slice(argv, argc)

	for _, val := range values {
		fmt.Printf("arg: %v\n", C.GoString(val))
	}

	payload := cgo.Handle(clientData).Value().(*command.CallbackPayload)
	fmt.Printf("unique data: %v\n", payload.UniqueData)

	payload.Callback(payload)

	return C.TCL_OK
}

func main() {
	interpreter := C.Tcl_CreateInterp()
	C.Tcl_Init(interpreter)

	name := C.CString("your_command")

	// Totally shit handling of C interop.
	// https://github.com/golang/go/issues/19837
	// https://github.com/golang/go/issues/19835
	fwdRef := (*[0]byte)(unsafe.Pointer(C.CommandWrapper))

	fn := func(*command.CallbackPayload) {
		fmt.Println("printing from inside the function.")
	}

	payload := &command.CallbackPayload{
		UniqueData: "yolo",
		Callback:   fn,
	}

	handle := C.uintptr_t(cgo.NewHandle(payload))

	ret := C.RegisterTclCommand(interpreter, name, fwdRef, handle)
	if ret != C.TCL_OK {
		fmt.Println("Failed to register Tcl command")
		return
	}

	code := `your_command arg1 arg2 arg3`
	retCode := C.Tcl_Eval(interpreter, C.CString(code))
	if retCode != C.TCL_OK {
		errMessage := C.GoString(C.Tcl_GetString(C.Tcl_GetObjResult(interpreter)))
		fmt.Println("Error:", errMessage)
	}

	// Clean up
	C.Tcl_DeleteInterp(interpreter)
}
