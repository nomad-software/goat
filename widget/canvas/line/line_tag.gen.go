// Code generated by tooling; DO NOT EDIT.
package line

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"

)



// SetTags sets the tags.
func (el *Line) SetTags(tags ...string) {
	tagStr := strings.Join(tags, " ")
	tk.Get().Eval("%s itemconfigure %s -tags [list %s]", el.GetParent().GetID(), el.GetID(), tagStr)
}
