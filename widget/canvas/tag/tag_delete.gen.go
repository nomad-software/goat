// Code generated by tooling; DO NOT EDIT.
package tag

import (
	"github.com/nomad-software/goat/internal/tk"

)



// Delete remove this item from the canvas.
func (el *Tag) Delete() {
	tk.Get().Eval("%s delete %s", el.GetParent().GetID(), el.GetID())
}
