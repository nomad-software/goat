// Code generated by tooling; DO NOT EDIT.
package entry

import (
	"github.com/nomad-software/goat/tk"
)




// AlightText aligns the text in different ways.
// See [widget.geometry.align]
func (el *Entry) AlignText(align string) {
	tk.Get().Eval("%s configure -justify {%s}", el.GetID(), align)
}