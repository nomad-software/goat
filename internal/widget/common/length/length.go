package length

import (
	"github.com/nomad-software/goat/internal/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetLength sets the length.
func (el stub) SetLength(l int) {
	tk.Get().Eval("%s configure -length %d", el.GetID(), l)
}
