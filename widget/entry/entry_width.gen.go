// Code generated by tooling; DO NOT EDIT.
package entry

import (
	"github.com/nomad-software/goat/tk"
)




// SetWidth sets the width.
func (el *Entry) SetWidth(w int) {
	tk.Get().Eval("%s configure -width %d", el.GetID(), w)
}
