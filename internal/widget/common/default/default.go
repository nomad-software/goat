package default_

import (
	"github.com/nomad-software/goat/internal/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetDefault sets this wiget as the default.
func (el stub) SetDefault() {
	tk.Get().Eval("%s configure -default active", el.GetID())
}
