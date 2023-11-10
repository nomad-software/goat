package radiobutton

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

type RadioButton struct {
	widget.Widget

	textVar  string
	valueVar string
}

// New creates a new radio button.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_radiobutton.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=command
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=image
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=invoke
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=stringvar -methods=GetValue,SetValue
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=textvar -methods=GetText,SetText
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=underline
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=width
func New(parent element.Element, text string) *RadioButton {
	button := &RadioButton{}
	button.SetParent(parent)
	button.SetType("radiobutton")

	button.textVar = variable.GenerateName(button.GetID())

	if parent != nil {
		button.valueVar = variable.GenerateName(button.GetType(), button.GetParent().GetID())
	} else {
		button.valueVar = variable.GenerateName(button.GetType())
	}

	tk.Get().Eval("ttk::radiobutton %s -textvariable %s -variable %s", button.GetID(), button.textVar, button.valueVar)

	button.SetText(text)

	return button
}

// Select selects the radio button.
func (el *RadioButton) Select() {
	tk.Get().SetVarStrValue(el.valueVar, "1")
}

// Deselect deselects the radio button.
func (el *RadioButton) Deselect() {
	tk.Get().SetVarStrValue(el.valueVar, "0")
}

// IsSelected return true if the radio button is selected.
func (el *RadioButton) IsSelected() bool {
	return tk.Get().GetVarBoolValue(el.valueVar)
}
