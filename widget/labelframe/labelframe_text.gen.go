// Code generated by tooling; DO NOT EDIT.
package labelframe

import (
	"github.com/nomad-software/goat/internal/tk"
)




// SetText sets the text.
func (el *LabelFrame) SetText(text string) {
	tk.Get().Eval("%s configure -text {%s}", el.GetID(), text)
}

// GetText gets the text.
func (el *LabelFrame) GetText() string {
	tk.Get().Eval("%s cget -text", el.GetID())
	return tk.Get().GetStrResult()
}
