package window

import (
	"github.com/nomad-software/goat/element"
	"github.com/nomad-software/goat/element/ui"
	"github.com/nomad-software/goat/tk"
)

// Window is the struct representing a window.
type Window struct {
	ui.Ele
}

// New creates a new window.
// The parent will usually be another window.
func New(parent element.Element) *Window {
	win := &Window{}
	win.SetParent(parent)
	win.SetType("window")

	// Create and show the window.
	tk.Get().Eval("toplevel %s", win.GetID())

	return win
}

// SetSize sets the window size.
func (w *Window) SetSize(width, height int) {
	tk.Get().Eval("wm geometry %s {%dx%d}", w.GetID(), width, height)
}

// SetGeometry sets the window size and position.
func (w *Window) SetGeometry(width, height, x, y int) {
	tk.Get().Eval("wm geometry %s {%dx%d+%d+%d}", w.GetID(), width, height, x, y)
}

// SetTitle sets the window title.
func (w *Window) SetTitle(title string) {
	tk.Get().Eval("wm title %s {%s}", w.GetID(), title)
}

// WaitForVisibility waits until this window is visible before continuing.
func (w *Window) WaitForVisibility() {
	tk.Get().Eval("tkwait visibility %s", w.GetID())
}

// SetOpacity sets the window opacity if it's supported.
func (w *Window) SetOpacity(opacity float64) {
	tk.Get().Eval("wm attributes %s -alpha %v", w.GetID(), opacity)
}

// SetFullScreen sets the window to be fullscreen or not.
func (w *Window) SetFullScreen(fullscreen bool) {
	tk.Get().Eval("wm attributes %s -fullscreen %v", w.GetID(), fullscreen)
}
