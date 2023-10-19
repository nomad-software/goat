package window

import "github.com/nomad-software/goat/element"

// Window is the struct representing a window.
type Window struct {
	element.ElementImpl
}

// New creates a new window.
// The parent will usually be another window.
func New(parent element.Element) *Window {
	win := &Window{}
	win.SetHash(win.GenerateHash())
	win.SetType("window")
	win.SetParent(parent)

	// Create and show the window.
	win.GetTk().Eval("toplevel %s", win.GetID())

	// Raise it above the parent.
	win.GetTk().Eval("raise %s %s", win.GetID(), parent.GetID())

	return win
}

// SetSize sets the window size.
func (w *Window) SetSize(width, height int) {
	w.GetTk().Eval("wm geometry %s %dx%d", w.GetID(), width, height)
}

// SetTitle sets the window title.
func (w *Window) SetTitle(title string) {
	w.GetTk().Eval("wm title %s {%s}", w.GetID(), title)
}

// WaitForVisibility waits until this window is visible before continuing.
func (w *Window) WaitForVisibility() {
	w.GetTk().Eval("tkwait visibility %s", w.GetID())
}
