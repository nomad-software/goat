package underline

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetPadding sets the padding.
func (el stub) SetUnderline(index int) {
	tk.Get().Eval("%s configure -underline %d", el.GetID(), index)
}
