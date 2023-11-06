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

	attached element.Element
}

// New creates a new x scrollbar.
func NewVertical(parent element.Element) *VerticalScrollbar {
	bar := &VerticalScrollbar{}
	bar.SetParent(parent)
	bar.SetType("yscrollbar")

	tk.Get().Eval("ttk::scrollbar %s -orient vertical", bar.GetID())

	return bar
}

// Set the attached widget.
func (s *VerticalScrollbar) AttachWidget(el element.Element) {
	s.attached = el

	tk.Get().Eval("%s configure -command [list %s yview]", s.GetID(), el.GetID())
}

// MoveTo moves the scrollbar so the widget should adjust its view so that the
// point given by fraction appears at the beginning of the widget. If fraction
// is 0 it refers to the beginning of the document. 1.0 refers to the end of
// the document, 0.333 refers to a point one-third of the way through the
// document, and so on.
func (s *VerticalScrollbar) MoveTo(fraction float64) {
	tk.Get().Eval("%s yview moveto {%v}", s.attached.GetID(), fraction)
}

// ScrollUnits moves the scrollbar so the widget should adjust its view by
// number units. The units are defined in whatever way makes sense for the
// widget, such as characters or lines in a text widget. Number is either 1,
// which means one unit should scroll off the top or left of the window, or -1,
// which means that one unit should scroll off the bottom or right of the
// window.
func (s *VerticalScrollbar) ScrollUnits(units int) {
	tk.Get().Eval("%s yview scroll %d units", s.attached.GetID(), units)
}

// ScrollPages moves the scrollbar so the widget should adjust its view by
// number pages. It is up to the widget to define the meaning of a page;
// typically it is slightly less than what fits in the window, so that there is
// a slight overlap between the old and new views. Number is either 1, which
// means the next page should become visible, or -1, which means that the
// previous page should become visible.
func (s *VerticalScrollbar) ScrollPages(pages int) {
	tk.Get().Eval("%s yview scroll %d pages", s.attached.GetID(), pages)
}
