// Code generated by tooling; DO NOT EDIT.
package checkbutton

import (
	"github.com/nomad-software/goat/internal/tk"

)







// SetValue sets the value.
func (el *CheckButton) SetValue(val string) {
	tk.Get().SetVarStrValue(el.valueVar, val)
}

// GetValue gets the value.
func (el *CheckButton) GetValue() string {
	return tk.Get().GetVarStrValue(el.valueVar)
}


