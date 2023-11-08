package selection

import (
	"github.com/nomad-software/goat/internal/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SelectText selects the text according to the passed limits.
func (el stub) SelectText(start, end int) {
	tk.Get().Eval("%s selection range %d %d", el.GetID(), start, end)
}

// IsTextSelected returns true if text is selected.
func (el stub) IsTextSelected() bool {
	tk.Get().Eval("%s selection present", el.GetID())
	return tk.Get().GetBoolResult()
}

// DeselectText deselected all text.
func (el stub) DeselectText() {
	tk.Get().Eval("%s selection clear", el.GetID())
}
