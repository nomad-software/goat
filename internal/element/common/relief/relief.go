package border

import (
	"github.com/nomad-software/goat/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetRelief sets the relief effect.
// See [element.relief]
func (el stub) SetRelief(r string) {
	tk.Get().Eval("%s configure -relief {%s}", el.GetID(), r)
}
