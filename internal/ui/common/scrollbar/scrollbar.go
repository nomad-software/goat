package scrollbar

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/ui/element" // IGNORE
	"github.com/nomad-software/goat/widget/scrollbar"
)

type stub struct{ element.Element } // IGNORE
func (el stub) GetID() string       { return "" } // IGNORE

// AttachHorizontalScrollbar sets the horizontal scrollbar.
func (el stub) AttachHorizontalScrollbar(bar *scrollbar.HorizontalScrollbar) {
	tk.Get().Eval("%s configure -xscrollcommand [list %s set]", el.GetID(), bar.GetID())
}

// AttachVerticalScrollbar sets the vertical scrollbar.
func (el stub) AttachVerticalScrollbar(bar *scrollbar.VerticalScrollbar) {
	tk.Get().Eval("%s configure -yscrollcommand [list %s set]", el.GetID(), bar.GetID())
}
