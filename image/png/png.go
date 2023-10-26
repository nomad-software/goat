package png

import (
	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/tk"
)

// Png is a png image.
type Png struct {
	image.Img
}

// New creates a png image.
func New(base64 string) *Png {
	img := &Png{}
	img.SetType("png")

	tk.Get().Eval("image create photo %s -format {png}", img.GetID())
	tk.Get().Eval("%s configure -data {%s}", img.GetID(), base64)

	return img
}
