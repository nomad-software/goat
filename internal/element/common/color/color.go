package color

import (
	"github.com/nomad-software/goat/tk"
)

type stub struct{}             // IGNORE
func (ele stub) GetID() string { return "" } // IGNORE

// SetForegroundColor sets the foreground color.
// See [element.color] for color names.
func (ele stub) SetForegroundColor(c string) {
	tk.Get().Eval("%s configure -foreground %s", ele.GetID(), c)
}

// SetBackgroundColor sets the background color.
// See [element.color] for color names.
func (ele stub) SetBackgroundColor(c string) {
	tk.Get().Eval("%s configure -background %s", ele.GetID(), c)
}

// SetInsertColor sets the insert color.
// See [element.color] for color names.
func (ele stub) SetInsertColor(c string) {
	tk.Get().Eval("%s configure -insertbackground %s", ele.GetID(), c)
}
