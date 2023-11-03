package border

import (
	"github.com/nomad-software/goat/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetBorderWidth sets the border width.
func (el stub) SetBorderWidth(b int) {
	tk.Get().Eval("%s configure -borderwidth %d", el.GetID(), b)
}
