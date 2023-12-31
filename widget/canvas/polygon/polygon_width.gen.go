// Code generated by tooling; DO NOT EDIT.
package polygon

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetOutlineWidth sets the outline width.
func (el *Polygon) SetOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -width %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetOutlineWidth sets the hover outline width.
func (el *Polygon) SetHoverOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -activewidth %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetOutlineWidth sets the disabled outline width.
func (el *Polygon) SetDisabledOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -disabledwidth %v", el.GetParent().GetID(), el.GetID(), width)
}
