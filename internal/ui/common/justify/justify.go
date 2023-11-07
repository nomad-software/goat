package justify

import (
	"github.com/nomad-software/goat/internal/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// AlightText aligns the text in different ways.
// See [option.align]
func (el stub) AlignText(align string) {
	tk.Get().Eval("%s configure -justify {%s}", el.GetID(), align)
}
