package ui

import (
	"strconv"
	"strings"

	"github.com/nomad-software/goat/internal/log"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/ui/element"
)

// Ele provides a base implementation of an ui element.
type Ele struct {
	element.Ele
}

// Update causes operations that are normally deferred, such as display updates
// and window layout calculations, to be performed immediately.
func (e *Ele) Update() {
	tk.Get().Eval("update idletasks")
}

// GetClass gets the ui element class.
// See [element.class] for class names.
func (e *Ele) GetClass() string {
	tk.Get().Eval("%s cget -class ", e.GetID())

	result := tk.Get().GetStrResult()

	if result == "" {
		tk.Get().Eval("winfo class %s", e.GetID())
		result = tk.Get().GetStrResult()
	}

	return result
}

// GetStyle gets the ui element class.
// See [element.style] for style names.
func (e *Ele) GetStyle() string {
	tk.Get().Eval("%s cget -style ", e.GetID())
	return tk.Get().GetStrResult()
}

// SetCursor sets the cursor of the ui element.
// See [option.cursor] for cursor names.
func (e *Ele) SetCursor(cursor string) {
	tk.Get().Eval("%s configure -cursor {%s}", e.GetID(), cursor)
}

// GetCursor gets the cursor of the ui element.
// See [option.cursor] for cursor names.
func (e *Ele) GetCursor() string {
	tk.Get().Eval("%s cget -cursor", e.GetID())
	return tk.Get().GetStrResult()
}

// SetKeyboadFocus sets that this ui element accepts the focus during keyboard
// traversal.
func (e *Ele) SetKeyboadFocus(focus bool) {
	tk.Get().Eval("%s configure -takefocus %v", e.GetID(), focus)
}

// AcceptsKeyboadFocus returns true if this ui element accepts the focus during
// keyboard traversal.
func (e *Ele) AcceptsKeyboadFocus() bool {
	tk.Get().Eval("%s cget -takefocus", e.GetID())
	return tk.Get().GetBoolResult()
}

// Destroy removes the ui element from the UI and cleans up its resources. Once
// destroyed you cannot refer to this ui element again or you will get a bad
// path name error from the interpreter.
func (e *Ele) Destroy() {
	tk.Get().Eval("destroy %s", e.GetID())
	e.SetType("destroyed")
}

// GetWidth gets the width of the ui element.
//
// Returns an int giving a ui element width in pixels. When a ui element is
// first created its width will be 1 pixel; the width will eventually be
// changed by a geometry manager to fulfil the window's needs.
func (e *Ele) GetWidth() int {
	tk.Get().Eval("winfo width %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetHeight gets the height of the ui element.
//
// Returns an int giving a ui element height in pixels. When a ui element
// is first created its height will be 1 pixel; the height will eventually be
// changed by a geometry manager to fulfil the window's needs.
func (e *Ele) GetHeight() int {
	tk.Get().Eval("winfo height %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetOSHandle gets the OS specific window handle.
//
// Returns a low-level platform-specific identifier for a window. On Unix
// platforms, this is the X window identifier. Under Windows, this is the
// Windows HWND. On the Macintosh the value has no meaning outside Tk.
func (e *Ele) GetOSHandle() int64 {
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
// list with two integers, which are the pointer's x and y coordinates measured
// in pixels in the screen's root window. If a virtual root window is in use on
// the screen, the position is computed in the virtual root. If the mouse
// pointer is not on the same screen as ui element then both of the returned
// coordinates are -1.
func (e *Ele) GetCursorPos() []int {
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
func (e *Ele) GetCursorXPos() int {
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
func (e *Ele) GetCursorYPos() int {
	tk.Get().Eval("winfo pointery %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetScreenWidth gets the width of the screen this ui element is on.
func (e *Ele) GetScreenWidth() int {
	tk.Get().Eval("winfo screenwidth %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetScreenHeight gets the height of the screen this ui element is on.
func (e *Ele) GetScreenHeight() int {
	tk.Get().Eval("winfo screenheight %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetXPos gets the x position of the ui element.
// You may need to wait until the ui element has been updated for this to
// return the correct value.
func (e *Ele) GetXPos(relativeToParent bool) int {
	if relativeToParent {
		tk.Get().Eval("winfo x %s", e.GetID())
	} else {
		tk.Get().Eval("winfo rootx %s", e.GetID())
	}
	return tk.Get().GetIntResult()
}

// GetYPos gets the y position of the ui element.
// You may need to wait until the ui element has been updated for this to
// return the correct value.
func (e *Ele) GetYPos(relativeToParent bool) int {
	if relativeToParent {
		tk.Get().Eval("winfo y %s", e.GetID())
	} else {
		tk.Get().Eval("winfo rooty %s", e.GetID())
	}
	return tk.Get().GetIntResult()
}

// Focus gives focus to the ui element.
func (e *Ele) Focus(force bool) {
	if force {
		tk.Get().Eval("focus -force %s", e.GetID())
	} else {
		tk.Get().Eval("focus %s", e.GetID())
	}
}

// Lower lowers a ui element below another if specified or below all of its
// siblings in the stacking order
func (e *Ele) Lower(el element.Element) {
	if el != nil {
		tk.Get().Eval("lower %s %s", e.GetID(), el.GetID())
	} else {
		tk.Get().Eval("lower %s", e.GetID())
	}
}

// Raise raises a ui element above another if specified or above all of its
// siblings in the stacking order.
func (e *Ele) Raise(el element.Element) {
	if el != nil {
		tk.Get().Eval("raise %s %s", e.GetID(), el.GetID())
	} else {
		tk.Get().Eval("raise %s", e.GetID())
	}
}