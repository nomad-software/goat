package window

import (
	"github.com/nomad-software/goat/element"
	"github.com/nomad-software/goat/element/ui"
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/tk/command"
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

// WaitForVisibility waits until this window is visible.
// This is typically used to wait for a newly-created window to appear on
// the screen before taking some action.
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

// SetTopmost sets the window to be the top-most. This makes the window not
// able to be lowered behind any others.
func (w *Window) SetTopmost(top bool) {
	tk.Get().Eval("wm attributes %s -topmost %v", w.GetID(), top)
}

// SetIconify sets whether the window is minimised.
func (w *Window) SetIconify(iconify bool) {
	if iconify {
		tk.Get().Eval("wm iconify %s", w.GetID())
	} else {
		tk.Get().Eval("wm deiconify %s", w.GetID())
	}
}

// SetMinSize sets the minimum size of the window.
func (w *Window) SetMinSize(width, height int) {
	tk.Get().Eval("wm minsize %s %d %d", w.GetID(), width, height)
}

// SetMinSize sets the maximum size of the window.
func (w *Window) SetMaxSize(width, height int) {
	tk.Get().Eval("wm maxsize %s %d %d", w.GetID(), width, height)
}

// BindProtocol binds a callback to be called when the specified protocol
// is triggered. A window manager protocol is a class of messages sent from a
// window manager to a Tk application outside of the normal event processing
// system.
// See [window.protocol] for protocol names.
func (w *Window) BindProtocol(protocol string, callback command.Callback) {
	name := command.GenerateName(protocol)
	tk.Get().CreateCommand(name, callback)
	tk.Get().Eval("wm protocol %s {%s} {%s}", w.GetID(), protocol, name)
}

// UnBindProtocol unbinds a previously bound callback.
func (w *Window) UnBindProtocol(protocol string) {
	name := command.GenerateName(protocol)
	tk.Get().Eval("wm protocol %s {%s} {}", w.GetID(), protocol)
	tk.Get().DeleteCommand(name)
}

// SetResizeable sets if a window width and height can be resized.
func (w *Window) SetResizeable(width, height bool) {
	tk.Get().Eval("wm resizable %s %v %v", w.GetID(), width, height)
}

// IsAbove returns if this window is above another.
func (w *Window) IsAbove(other *Window) bool {
	tk.Get().Eval("wm stackorder %s isabove %s", w.GetID(), other.GetID())
	return tk.Get().GetBoolResult()
}

// IsAbove returns if this window is above another.
func (w *Window) IsBelow(other *Window) bool {
	tk.Get().Eval("wm stackorder %s isbelow %s", w.GetID(), other.GetID())
	return tk.Get().GetBoolResult()
}

// Wait waits for the window to be destroyed.
// This is typically used to wait for a user to finish interacting with a
// dialog box before using the result of that interaction.
func (w *Window) Wait() {
	tk.Get().Eval("tkwait window %s", w.GetID())
}
