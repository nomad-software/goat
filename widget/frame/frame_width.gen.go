// Code generated by tooling; DO NOT EDIT.
package frame

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetWidth sets the width.
func (el *Frame) SetWidth(w int) {
	tk.Get().Eval("%s configure -width %d", el.GetID(), w)
}
