// Code generated by tooling; DO NOT EDIT.
package widget

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetState sets the state.
// See [option.state] for state values.
func (el *Widget) SetState(state string) {
	tk.Get().Eval("%s itemconfigure %s -state {%s}", el.GetParent().GetID(), el.GetID(), state)
}
