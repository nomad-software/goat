// Code generated by tooling; DO NOT EDIT.
package frame

import (
	"github.com/nomad-software/goat/internal/tk"
)




// SetPadding sets the padding.
func (el *Frame) SetPadding(p int) {
	tk.Get().Eval("%s configure -padding %d", el.GetID(), p)
}
