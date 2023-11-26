package dash

import (
	"fmt"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetDash sets the outline dash.
// Each element represents the number of pixels of a line segment. Only the odd
// segments are drawn using the “outline” color. The other segments are drawn
// transparent.
func (el stub) SetOutlineDash(first, second float64, others ...float64) {
	otherStr := ""
	for _, i := range others {
		otherStr += fmt.Sprintf(" %v", i)
	}
	tk.Get().Eval("%s itemconfigure %s -dash [list %v %v %s]", el.GetParent().GetID(), el.GetID(), first, second, otherStr)
}

// SetActiveDash sets the active outline dash.
func (el stub) SetOutlineActiveDash(first, second float64, others ...float64) {
	otherStr := ""
	for _, i := range others {
		otherStr += fmt.Sprintf(" %v", i)
	}
	tk.Get().Eval("%s itemconfigure %s -activedash [list %v %v %s]", el.GetParent().GetID(), el.GetID(), first, second, otherStr)
}

// SetDisabledDash sets the disabled outline dash.
func (el stub) SetOutlineDisabledDash(first, second float64, others ...float64) {
	otherStr := ""
	for _, i := range others {
		otherStr += fmt.Sprintf(" %v", i)
	}
	tk.Get().Eval("%s itemconfigure %s -disableddash [list %v %v %s]", el.GetParent().GetID(), el.GetID(), first, second, otherStr)
}

// SetDashOffset sets the starting offset in pixels.
func (el stub) SetOutlineDashOffset(offset float64) {
	tk.Get().Eval("%s itemconfigure %s -dashoffset %v", el.GetParent().GetID(), el.GetID(), offset)
}
