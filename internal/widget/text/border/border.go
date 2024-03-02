package border

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetBorderWidth sets border width.
func (el stub) SetBorderWidth(width int) {
	tk.Get().Eval("%s tag configure {%s} -borderwidth %d", el.GetParent().GetID(), el.GetID(), width)
}

// SetRelief sets the relief effect.
// See [option.relief] for relief values.
func (el stub) SetRelief(relief string) {
	tk.Get().Eval("%s tag configure {%s} -relief {%s}", el.GetParent().GetID(), el.GetID(), relief)
}
