package scrollbar

import (
	"github.com/nomad-software/goat/internal/element"
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/widget"
)

// A horizontal scrollbar used to scroll the content of widgets.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_scrollbar.html
type HorizontalScrollbar struct {
	widget.Widget
}

// New creates a new x scrollbar.
func NewHorizontal(parent element.Element) *HorizontalScrollbar {
	bar := &HorizontalScrollbar{}
	bar.SetParent(parent)
	bar.SetType("xscrollbar")

	tk.Get().Eval("ttk::scrollbar %s -orient horizontal", bar.GetID())

	return bar
}
