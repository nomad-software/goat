// Code generated by tooling; DO NOT EDIT.
package radiobutton

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetPadding sets the padding.
func (el *RadioButton) SetUnderline(index int) {
	tk.Get().Eval("%s configure -underline %d", el.GetID(), index)
}
