package fill

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetFill sets the fill color.
// See [option.color] for color names.
func (el stub) SetFillColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -fill {%s}", el.GetParent().GetID(), el.GetID(), color)
}

// SetActiveFill sets the active fill color.
// See [option.color] for color names.
func (el stub) SetActiveFillColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -activefill {%s}", el.GetParent().GetID(), el.GetID(), color)
}

// SetDisabledFill sets the active fill color.
// See [option.color] for color names.
func (el stub) SetDisabledFillColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -disabledfill {%s}", el.GetParent().GetID(), el.GetID(), color)
}
