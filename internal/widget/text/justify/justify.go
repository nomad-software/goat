package justify

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// AlightText aligns the text in different ways.
// See [option.align]
func (el stub) AlignText(align string) {
	tk.Get().Eval("%s tag configure {%s} -justify {%s}", el.GetParent().GetID(), el.GetID(), align)
}
