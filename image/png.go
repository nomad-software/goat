package image

import (
	"github.com/nomad-software/goat/element"
	"github.com/nomad-software/goat/tk"
)

// Png is an image.
type Png struct {
	element.Ele
}

// New creates an image.
func New(base64 string) *Png {
	img := &Png{}
	img.SetType("image")

	tk.Get().Eval("image create photo %s", img.GetID())
	tk.Get().Eval("%s configure -data {%s}", img.GetID(), base64)

	return img
}
