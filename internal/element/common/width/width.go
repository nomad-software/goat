package width

import (
	"github.com/nomad-software/goat/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetWidth sets the width.
func (el stub) SetWidth(w int) {
	tk.Get().Eval("%s configure -width %d", el.GetID(), w)
}
