package scrollbar

import (
	"github.com/nomad-software/goat/internal/element" // IGNORE
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/widget/scrollbar"
)

type stub struct{ element.Element } // IGNORE
func (el stub) GetID() string       { return "" } // IGNORE

// SetHorizontalScrollbar sets the horizontal scrollbar.
func (el stub) SetHorizontalScrollbar(bar *scrollbar.HorizontalScrollbar) {
	tk.Get().Eval("%s configure -command [list %s xview]", bar.GetID(), el.GetID())
	tk.Get().Eval("%s configure -xscrollcommand [list %s set]", el.GetID(), bar.GetID())
}

// SetVerticalScrollbar sets the vertical scrollbar.
func (el stub) SetVerticalScrollbar(bar *scrollbar.VerticalScrollbar) {
	tk.Get().Eval("%s configure -command [list %s yview]", bar.GetID(), el.GetID())
	tk.Get().Eval("%s configure -yscrollcommand [list %s set]", el.GetID(), bar.GetID())
}
