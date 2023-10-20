package window

import (
	"github.com/nomad-software/goat/element"
	"github.com/nomad-software/goat/tk"
)

// Window is the struct representing a window.
type Window struct {
	element.ElementImpl
}

// New creates a new window.
// The parent will usually be another window.
func New(parent element.Element) *Window {
	win := &Window{}
	win.SetParent(parent)
	win.SetType("window")

	// Create and show the window.
	tk.Get().Eval("toplevel %s", win.GetID())

	// Raise it above the parent.
	tk.Get().Eval("raise %s %s", win.GetID(), parent.GetID())

	return win
}

// SetSize sets the window size.
func (w *Window) SetSize(width, height int) {
	tk.Get().Eval("wm geometry %s %dx%d", w.GetID(), width, height)
}

// SetTitle sets the window title.
func (w *Window) SetTitle(title string) {
	tk.Get().Eval("wm title %s {%s}", w.GetID(), title)
}

// WaitForVisibility waits until this window is visible before continuing.
func (w *Window) WaitForVisibility() {
	tk.Get().Eval("tkwait visibility %s", w.GetID())
}
