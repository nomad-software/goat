// Code generated by tooling; DO NOT EDIT.
package checkbutton

import (
	"github.com/nomad-software/goat/internal/tk"

)



// Invoke invokes the command associated with this widget.
func (el *CheckButton) Invoke() {
	tk.Get().Eval("%s invoke", el.GetID())
}
