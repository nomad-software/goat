package listview

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

// Row represents a row in the list view.
type Row struct {
	element.Ele
}

// GetValues gets the rows values.
func (el *Row) GetValues() []string {
	tk.Get().Eval("%s item %s -values", el.GetParent().GetID(), el.GetID())
	return tk.Get().GetStrSliceResult()
}
