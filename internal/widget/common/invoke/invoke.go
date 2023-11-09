package invoke

import (
	"github.com/nomad-software/goat/internal/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// Invoke invokes the command associated with this widget.
func (el stub) Invoke() {
	tk.Get().Eval("%s invoke", el.GetID())
}
