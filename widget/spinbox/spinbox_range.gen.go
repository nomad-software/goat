// Code generated by tooling; DO NOT EDIT.
package spinbox

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetFromValue sets the from value of a range.
func (el *Spinbox) SetFromValue(val float64) {
	tk.Get().Eval("%s configure -from {%v}", el.GetID(), val)
}

// SetToValue sets the to value of a range.
func (el *Spinbox) SetToValue(val float64) {
	tk.Get().Eval("%s configure -to {%v}", el.GetID(), val)
}
