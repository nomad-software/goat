// Code generated by tooling; DO NOT EDIT.
package notebook

import (
	"github.com/nomad-software/goat/internal/tk"
)




// SetWidth sets the width.
func (el *NoteBook) SetWidth(w int) {
	tk.Get().Eval("%s configure -width %d", el.GetID(), w)
}
