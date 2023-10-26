package gif

import (
	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/tk"
)

// Gif is a gif image..
type Gif struct {
	image.Img
}

// New creates a gif image.
func New(base64 string) *Gif {
	img := &Gif{}
	img.SetType("gif")

	tk.Get().Eval("image create photo %s -format {gif}", img.GetID())
	tk.Get().Eval("%s configure -data {%s}", img.GetID(), base64)

	return img
}
