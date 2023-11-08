package delete_

import (
	"github.com/nomad-software/goat/internal/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// DeleteText deletes all text.
func (el stub) DeleteText(start, end int) {
	tk.Get().Eval("%s delete %d %d", el.GetID(), start, end)
}
