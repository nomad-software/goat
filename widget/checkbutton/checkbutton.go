package checkbutton

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

type CheckButton struct {
	widget.Widget

	textVar  string
	valueVar string
}

// New creates a new checkbutton.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_checkbutton.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=command
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=image
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=invoke
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=stringvar -methods=GetValue,SetValue
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=textvar -methods=GetText,SetText
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=underline
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=width
func New(parent element.Element, text string) *CheckButton {
	button := &CheckButton{}
	button.SetParent(parent)
	button.SetType("checkbutton")

	button.textVar = variable.GenerateName("textvar", button.GetID())
	button.valueVar = variable.GenerateName("valuevar", button.GetID())

	tk.Get().Eval("ttk::checkbutton %s -textvariable %s -variable %s", button.GetID(), button.textVar, button.valueVar)

	button.SetText(text)
	button.SetValue("0")

	return button
}

// Check checks the checkbutton.
func (el *CheckButton) Check() {
	tk.Get().SetVarStrValue(el.valueVar, "1")
}

// Check unchecks the checkbutton.
func (el *CheckButton) UnCheck() {
	tk.Get().SetVarStrValue(el.valueVar, "0")
}

// Check half-checks the checkbutton.
func (el *CheckButton) HalfCheck() {
	tk.Get().SetVarStrValue(el.valueVar, "")
}

// IsChecked returns true if the checkbutton is checked.
func (el *CheckButton) IsChecked() bool {
	return tk.Get().GetVarBoolValue(el.valueVar)
}

// Destroy removes the ui element from the UI and cleans up its resources. Once
// destroyed you cannot refer to this ui element again or you will get a bad
// path name error from the interpreter.
func (el *CheckButton) Destroy() {
	el.Ele.Destroy()
	tk.Get().DeleteVar(el.textVar)
	tk.Get().DeleteVar(el.valueVar)
}
