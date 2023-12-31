// Code generated by tooling; DO NOT EDIT.
package line

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetOutlineWidth sets the outline width.
func (el *Line) SetOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -width %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetOutlineWidth sets the hover outline width.
func (el *Line) SetHoverOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -activewidth %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetOutlineWidth sets the disabled outline width.
func (el *Line) SetDisabledOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -disabledwidth %v", el.GetParent().GetID(), el.GetID(), width)
}
