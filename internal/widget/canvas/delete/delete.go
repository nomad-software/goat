package delete

import (
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

type stub struct{ element.Element } // IGNORE

type deleter interface {
	DeleteItem(item element.Element)
}

// Delete removes this item from the canvas.
func (el stub) Delete() {
	el.GetParent().(deleter).DeleteItem(el)
}
