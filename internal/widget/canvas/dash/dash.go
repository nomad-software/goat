package dash

import (
	"fmt"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetDash sets the outline dash.
func (el stub) SetDash(one, two int, rest ...int) {
	str := ""
	for _, i := range rest {
		str += fmt.Sprintf(" %d", i)
	}
	tk.Get().Eval("%s itemconfigure %s -dash [list %d %d %s]", el.GetParent().GetID(), el.GetID(), one, two, str)
}
