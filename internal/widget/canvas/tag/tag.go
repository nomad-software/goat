package tag

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetTags sets the tags.
func (el stub) SetTags(tags ...string) {
	tagStr := strings.Join(tags, " ")
	tk.Get().Eval("%s itemconfigure %s -tags [list %s]", el.GetParent().GetID(), el.GetID(), tagStr)
}
