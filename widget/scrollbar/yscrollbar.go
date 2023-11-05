package scrollbar

import (
	"github.com/nomad-software/goat/internal/element"
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/widget"
)

// A vertical scrollbar used to scroll the content of widgets.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_scrollbar.html
type VerticalScrollbar struct {
	widget.Widget
}

// New creates a new x scrollbar.
func NewVertical(parent element.Element) *VerticalScrollbar {
	bar := &VerticalScrollbar{}
	bar.SetParent(parent)
	bar.SetType("yscrollbar")

	tk.Get().Eval("ttk::scrollbar %s -orient vertical", bar.GetID())

	return bar
}
