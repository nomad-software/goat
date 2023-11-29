package widget

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvaswidget"
)

// Widget represents a widget in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Widget -pkg=canvas/anchor
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Widget -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Widget -pkg=canvas/tag
type Widget struct {
	element.Ele
}

// SetWidth sets the width
func (el *Widget) SetWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -width %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetHeight sets the width
func (el *Widget) SetHeight(height float64) {
	tk.Get().Eval("%s itemconfigure %s -height %v", el.GetParent().GetID(), el.GetID(), height)
}

// SetWidget sets the widget.
func (el *Widget) SetWidget(e element.Element) {
	tk.Get().Eval("%s itemconfigure %s -window %s", el.GetParent().GetID(), el.GetID(), e.GetID())
}
