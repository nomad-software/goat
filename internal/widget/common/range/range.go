package range_

import (
	"github.com/nomad-software/goat/internal/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetFromValue sets the from value of a range.
func (el stub) SetFromValue(value float64) {
	tk.Get().Eval("%s configure -from {%v}", el.GetID(), value)
}

// SetToValue sets the to value of a range.
func (el stub) SetToValue(value float64) {
	tk.Get().Eval("%s configure -to {%v}", el.GetID(), value)
}
