package text

import (
	"github.com/nomad-software/goat/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetText sets the text.
func (el stub) SetText(text string) {
	tk.Get().Eval("%s configure -text {%s}", el.GetID(), text)
}

// GetText gets the text.
func (el stub) GetText() string {
	tk.Get().Eval("%s cget -text", el.GetID())
	return tk.Get().GetStrResult()
}
