package element

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/tk/command"
)

type UIElement interface {
	GetClass() string
	SetCursor(string)
	GetCursor() string
	// Bind(???)
	// UnBind(???)
	Destroy()
	GetWidth() int
	GetHeight() int
	GetOSHandle() int
	GetCursorPos() []int
	GetCursorXPos() int
	GetCursorYPos() int
	GetScreenWidth() int
	GetScreenHeight() int
	GetXPos(bool) int
	GetYPos(bool) int
	// GenerateEvent(string)
	Focus(bool)
	Lower(Element)
	Raise(Element)
}

// UIEle provides a base implementation of an ui element.
type UIEle struct {
	Ele
}

// GetClass gets the ui element class.
// See [class] package for class names.
func (e *UIEle) GetClass() string {
	tk.Get().Eval("%s cget -class ", e.GetID())

	result := tk.Get().GetStrResult()

	if result == "" {
		tk.Get().Eval("winfo class %s", e.GetID())
		result = tk.Get().GetStrResult()
	}

	return result
}

// SetCursor sets the cursor of the ui element.
// See [cursor] package for cursor names.
func (e *UIEle) SetCursor(cursor string) {
	tk.Get().Eval("%s configure -cursor {%s}", e.GetID(), cursor)
}

// GetCursor gets the cursor of the ui element.
// See [cursor] package for cursor names.
func (e *UIEle) GetCursor() string {
	tk.Get().Eval("%s cget -cursor", e.GetID())
	return tk.Get().GetStrResult()
}

// Destroy removes the element from the UI and cleans up its resources.
// Once destroyed you cannot refer to this ui element again or you will get a
// bad path name error from the interpreter.
func (e *UIEle) Destroy() {
	tk.Get().Eval("destroy %s", e.GetID())
	e.SetType("destroyed")
}

// GetWidth gets the width of the ui element.
//
// Returns an int giving a ui element width in pixels. When a ui element is
// first created its width will be 1 pixel; the width will eventually be
// changed by a geometry manager to fulfil the window's needs.
func (e *UIEle) GetWidth() int {
	tk.Get().Eval("winfo width %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetHeight gets the height of the ui element.
//
// Returns an int giving a ui element height in pixels. When a ui element
// is first created its height will be 1 pixel; the height will eventually be
// changed by a geometry manager to fulfil the window's needs.
func (e *UIEle) GetHeight() int {
	tk.Get().Eval("winfo height %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetOSHandle gets the OS specific window handle.
//
// Returns a low-level platform-specific identifier for a window. On Unix
// platforms, this is the X window identifier. Under Windows, this is the
// Windows HWND. On the Macintosh the value has no meaning outside Tk.
func (e *UIEle) GetOSHandle() int64 {
	tk.Get().Eval("winfo id %s", e.GetID())
	result := tk.Get().GetStrResult()

	// Remove the 0x prefix.
	if len(result) > 2 {
		result = result[2:]
	}

	hwnd, err := strconv.ParseInt(result, 16, 0)
	if err != nil {
		log.Error(err)
	}

	return hwnd
}

// GetCursorPos gets the x and y position of the cursor on the ui element.
//
// If the mouse pointer is on the same screen as the ui element, it returns a
// list with two elements, which are the pointer's x and y coordinates measured
// in pixels in the screen's root window. If a virtual root window is in use on
// the screen, the position is computed in the virtual root. If the mouse
// pointer is not on the same screen as ui element then both of the returned
// coordinates are -1.
func (e *UIEle) GetCursorPos() []int {
	tk.Get().Eval("winfo pointerxy %s", e.GetID())
	result := tk.Get().GetStrResult()

	strs := strings.Split(result, " ")
	pos := make([]int, 0)

	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Error(err)
		}
		pos = append(pos, i)
	}

	return pos
}

// GetCursorXPos gets the x position of the cursor on the ui element.
//
// If the mouse pointer is on the same screen as the ui element, it returns the
// pointer's x coordinate, measured in pixels in the screen's root window. If a
// virtual root window is in use on the screen, the position is measured in the
// virtual root. If the mouse pointer is not on the same screen as ui element
// then -1 is returned.
func (e *UIEle) GetCursorXPos() int {
	tk.Get().Eval("winfo pointerx %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetCursorYPos gets the y position of the cursor on the ui element.
//
// If the mouse pointer is on the same screen as the ui element, it returns the
// pointer's y coordinate, measured in pixels in the screen's root window. If a
// virtual root window is in use on the screen, the position is measured in the
// virtual root. If the mouse pointer is not on the same screen as ui element
// then -1 is returned.
func (e *UIEle) GetCursorYPos() int {
	tk.Get().Eval("winfo pointery %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetScreenWidth gets the width of the screen this ui element is on.
func (e *UIEle) GetScreenWidth() int {
	tk.Get().Eval("winfo screenwidth %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetScreenHeight gets the height of the screen this ui element is on.
func (e *UIEle) GetScreenHeight() int {
	tk.Get().Eval("winfo screenheight %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetXPos gets the x position of the ui element.
func (e *UIEle) GetXPos(relativeToParent bool) int {
	if relativeToParent {
		tk.Get().Eval("winfo x %s", e.GetID())
	} else {
		tk.Get().Eval("winfo rootx %s", e.GetID())
	}
	return tk.Get().GetIntResult()
}

// GetYPos gets the y position of the ui element.
func (e *UIEle) GetYPos(relativeToParent bool) int {
	if relativeToParent {
		tk.Get().Eval("winfo y %s", e.GetID())
	} else {
		tk.Get().Eval("winfo rooty %s", e.GetID())
	}
	return tk.Get().GetIntResult()
}

// Focus gives focus to the ui element.
func (e *UIEle) Focus(force bool) {
	if force {
		tk.Get().Eval("focus -force %s", e.GetID())
	} else {
		tk.Get().Eval("focus %s", e.GetID())
	}
}

// Lower lowers a ui element below another if specified or below all of its
// siblings in the stacking order
func (e *UIEle) Lower(el Element) {
	if el != nil {
		tk.Get().Eval("lower %s %s", e.GetID(), el.GetID())
	} else {
		tk.Get().Eval("lower %s", e.GetID())
	}
}

// Raise raises a ui element above another if specified or above all of its
// siblings in the stacking order.
func (e *UIEle) Raise(el Element) {
	if el != nil {
		tk.Get().Eval("raise %s %s", e.GetID(), el.GetID())
	} else {
		tk.Get().Eval("raise %s", e.GetID())
	}
}

func (e *UIEle) Bind(binding string, fn command.Callback) {
	name := command.GenerateName(binding, e.GetID())
	// proc, deleteProc := command.CreateProcedure(fn)

	// tk.Get.CreateCommand(name, proc, data, deleteProc)

	fmt.Printf("%v\n", name)
}
