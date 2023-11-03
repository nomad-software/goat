package height

import (
	"github.com/nomad-software/goat/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetHeight sets the height.
func (el stub) SetHeight(h int) {
	tk.Get().Eval("%s configure -height %d", el.GetID(), h)
}
