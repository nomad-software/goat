// Code generated by tooling; DO NOT EDIT.
package arc

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetFill sets the fill color.
// See [option.color] for color names.
func (el *Arc) SetFillColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -fill {%s}", el.GetParent().GetID(), el.GetID(), color)
}

// SetActiveFill sets the active fill color.
// See [option.color] for color names.
func (el *Arc) SetActiveFillColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -activefill {%s}", el.GetParent().GetID(), el.GetID(), color)
}

// SetDisabledFill sets the active fill color.
// See [option.color] for color names.
func (el *Arc) SetDisabledFillColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -disabledfill {%s}", el.GetParent().GetID(), el.GetID(), color)
}
