package arc

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

// Arc represents an arc in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/dash
type Arc struct {
	element.Ele
}

// SetStyle sets the style of arc.
// See [widget.canvas.arc.style] for style names.
func (el *Arc) SetStyle(style string) {
	tk.Get().Eval("%s itemconfigure %s -style {%s}", el.GetParent().GetID(), el.GetID(), style)
}
