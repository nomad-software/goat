// Code generated by tooling; DO NOT EDIT.
package entry

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SelectText selects the text according to the passed limits.
func (el *Entry) SelectText(start, end int) {
	tk.Get().Eval("%s selection range %d %d", el.GetID(), start, end)
}

// IsTextSelected returns true if text is selected.
func (el *Entry) IsTextSelected() bool {
	tk.Get().Eval("%s selection present", el.GetID())
	return tk.Get().GetBoolResult()
}

// DeselectText deselected all text.
func (el *Entry) DeselectText() {
	tk.Get().Eval("%s selection clear", el.GetID())
}
