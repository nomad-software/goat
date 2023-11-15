package data

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetData sets the data of a widget.
func (el stub) SetData(data ...string) {
	values := strings.Join(data, "} {")
	tk.Get().Eval("%s configure -values [list {%s}]", el.GetID(), values)
}

// GetData gets the set data of the widget.
// If no data has been set, this will return empty.
func (el stub) GetData() []string {
	tk.Get().Eval("%s cget -values", el.GetID())
	data := make([]string, 0)

	result := tk.Get().GetStrResult()
	if result != "" {
		values := strings.Split(result, " ")
		for _, val := range values {
			data = append(data, val[1:len(val)-1])
		}
	}

	return data
}
