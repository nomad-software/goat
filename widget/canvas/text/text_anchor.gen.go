// Code generated by tooling; DO NOT EDIT.
package text

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetAnchor sets the anchor.
// See [option.anchor] for anchor values.
func (el *Text) SetAnchor(anchor string) {
	tk.Get().Eval("%s itemconfigure %s -anchor {%s}", el.GetParent().GetID(), el.GetID(), anchor)
}
