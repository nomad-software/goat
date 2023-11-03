package padding

import (
	"github.com/nomad-software/goat/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetPadding sets the padding.
func (el stub) SetPadding(p int) {
	tk.Get().Eval("%s configure -padding %d", el.GetID(), p)
}
