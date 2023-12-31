// Code generated by tooling; DO NOT EDIT.
package combobox

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"

)



// SetData sets the data of a widget.
func (el *Combobox) SetData(data ...string) {
	values := strings.Join(data, "} {")
	tk.Get().Eval("%s configure -values [list {%s}]", el.GetID(), values)
}

// GetData gets the set data of the widget.
// If no data has been set, this will return empty.
func (el *Combobox) GetData() []string {
	tk.Get().Eval("%s cget -values", el.GetID())
	return tk.Get().GetStrSliceResult()
}
