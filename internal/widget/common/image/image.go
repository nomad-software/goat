package image

import (
	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetImage sets an image and its position.
// See [option.compound] for compound values.
func (el stub) SetImage(img *image.Image, compound string) {
	tk.Get().Eval("%s configure -image %s -compound {%s}", el.GetID(), img.GetID(), compound)
}
