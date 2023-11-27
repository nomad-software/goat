package outline

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetOutlineColor sets the outline color.
// See [option.color] for color names.
func (el stub) SetOutlineColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -outline {%s}", el.GetParent().GetID(), el.GetID(), color)
}

// SetHoverOutlineColor sets the hover outline color.
// See [option.color] for color names.
func (el stub) SetHoverOutlineColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -activeoutline {%s}", el.GetParent().GetID(), el.GetID(), color)
}

// SetDisabledOutlineColor sets the disabled outline color.
// See [option.color] for color names.
func (el stub) SetDisabledOutlineColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -disabledoutline {%s}", el.GetParent().GetID(), el.GetID(), color)
}

// SetOutlineWidth sets the outline width.
func (el stub) SetOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -width %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetOutlineWidth sets the hover outline width.
func (el stub) SetHoverOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -activewidth %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetOutlineWidth sets the disabled outline width.
func (el stub) SetDisabledOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -disabledwidth %v", el.GetParent().GetID(), el.GetID(), width)
}
