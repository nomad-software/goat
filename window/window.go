package window

import (
	"github.com/nomad-software/goat/element"
	"github.com/nomad-software/goat/tk"
)

// init configures the environment.
func init() {
	tk.Get().Eval("encoding system utf-8")
}

// MainWindow is the struct representing the main window.
type MainWindow struct {
	Window
}

// GetMain gets the main window of the application.
func GetMain() *MainWindow {
	win := &MainWindow{}
	win.SetID(".")
	win.SetType("window")

	return win
}

// Show shows the main window and starts the application.
// This method should not be deferred in the main function or else it will
// potentially trap panics in other parts of the program.
func (w *MainWindow) Show() {
	tk.Get().Start()
}

// Window is the struct representing a window.
type Window struct {
	element.UIElementImpl
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
	tk.Get().Eval("wm geometry %s {%dx%d}", w.GetID(), width, height)
}

// SetGeometry sets the window size.
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
