package underline

import (
	"github.com/nomad-software/goat/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetPadding sets the padding.
func (el stub) SetUnderline(index int) {
	tk.Get().Eval("%s configure -underline %d", el.GetID(), index)
}
