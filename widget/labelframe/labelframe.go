package labelframe

import (
	"github.com/nomad-software/goat/internal/element"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/widget"
)

// LabelFrame is a container used to group other widgets together. It has an
// optional label, which may be a plain text string or another widget.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_labelframe.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*LabelFrame -pkg=height
//go:generate go run ../../internal/tools/generate/main.go -recv=*LabelFrame -pkg=padding
//go:generate go run ../../internal/tools/generate/main.go -recv=*LabelFrame -pkg=text
//go:generate go run ../../internal/tools/generate/main.go -recv=*LabelFrame -pkg=underline
//go:generate go run ../../internal/tools/generate/main.go -recv=*LabelFrame -pkg=width
type LabelFrame struct {
	widget.Widget
}

// New creates a new label frame.
func New(parent element.Element, text string, underline int) *LabelFrame {
	frame := &LabelFrame{}
	frame.SetParent(parent)
	frame.SetType("labelframe")

	tk.Get().Eval("ttk::labelframe %s -text {%s} -underline %d", frame.GetID(), text, underline)

	return frame
}

// SetLabelAnchor sets the anchor for the label.
// See [option.anchor] for anchor values.
func (l *LabelFrame) SetLabelAnchor(anchor string) {
	tk.Get().Eval("%s configure -labelanchor {%s}", l.GetID(), anchor)
}
