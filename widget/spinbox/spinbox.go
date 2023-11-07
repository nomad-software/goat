package spinbox

import (
	"math"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/ui/element"
	"github.com/nomad-software/goat/widget"
)

type Spinbox struct {
	widget.Widget

	valueVar string
}

// New creates a new spinbox.
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=color -methods=SetForegroundColor
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=command
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=data
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=float
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=font
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=justify
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=range
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=scrollbar -methods=AttachHorizontalScrollbar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=show
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=width
func New(parent element.Element) *Spinbox {
	spinbox := &Spinbox{}
	spinbox.SetParent(parent)
	spinbox.SetType("spinbox")

	spinbox.valueVar = variable.GenerateName(spinbox.GetID())

	tk.Get().Eval("ttk::spinbox %s -textvariable %s", spinbox.GetID(), spinbox.valueVar)

	spinbox.SetFromValue(math.MinInt8)
	spinbox.SetToValue(math.MaxInt8)
	spinbox.SetValue(0)

	return spinbox
}

// SetStep sets the increment of each change.
func (el *Spinbox) SetStep(step float64) {
	tk.Get().Eval("%s configure -increment {%v}", el.GetID(), step)
}

// SetWrap sets if the value should wrap if it reaches the end.
func (el *Spinbox) SetWrap(wrap bool) {
	tk.Get().Eval("%s configure -wrap {%v}", el.GetID(), wrap)
}

// SetFormat sets the display format of the number.
// Before is the number of digits before a decimal place.
// After is the number of digits after a decimal place.
func (el *Spinbox) SetFormat(before, after int) {
	tk.Get().Eval("%s configure -format %%%d.%df", el.GetID(), before, after)
}
