// Code generated by tooling; DO NOT EDIT.
package window

import (
	"github.com/nomad-software/goat/internal/tk"
)




// SetWidth sets the width.
func (el *Window) SetWidth(w int) {
	tk.Get().Eval("%s configure -width %d", el.GetID(), w)
}
