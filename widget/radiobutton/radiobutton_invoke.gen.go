// Code generated by tooling; DO NOT EDIT.
package radiobutton

import (
	"github.com/nomad-software/goat/internal/tk"
)




// Invoke invokes the command associated with this widget.
func (el *RadioButton) Invoke() {
	tk.Get().Eval("%s invoke", el.GetID())
}