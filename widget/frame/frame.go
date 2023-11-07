package frame

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/ui/element"
	"github.com/nomad-software/goat/widget"
)

// A frame widget is a container, used to group other widgets together.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_frame.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Frame -pkg=borderwidth
//go:generate go run ../../internal/tools/generate/main.go -recv=*Frame -pkg=height
//go:generate go run ../../internal/tools/generate/main.go -recv=*Frame -pkg=padding
//go:generate go run ../../internal/tools/generate/main.go -recv=*Frame -pkg=padding
//go:generate go run ../../internal/tools/generate/main.go -recv=*Frame -pkg=width
type Frame struct {
	widget.Widget
}

// New creates a new frame.
func New(parent element.Element, borderWidth int, relief string) *Frame {
	frame := &Frame{}
	frame.SetParent(parent)
	frame.SetType("frame")

	tk.Get().Eval("ttk::frame %s -borderwidth %d -relief {%s}", frame.GetID(), borderWidth, relief)

	return frame
}
