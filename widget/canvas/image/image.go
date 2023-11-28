package image

import (
	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvasimage"
)

// Image represents an image in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Image -pkg=canvas/anchor
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Image -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Image -pkg=canvas/tag
type Image struct {
	element.Ele
}

// SetImage sets the image.
func (el *Image) SetImage(img *image.Image) {
	tk.Get().Eval("%s itemconfigure %s -image %v", el.GetParent().GetID(), el.GetID(), img.GetID())
}

// SetImage sets the hover image.
func (el *Image) SetHoverImage(img *image.Image) {
	tk.Get().Eval("%s itemconfigure %s -activeimage %v", el.GetParent().GetID(), el.GetID(), img.GetID())
}

// SetImage sets the disabled image.
func (el *Image) SetDisabledImage(img *image.Image) {
	tk.Get().Eval("%s itemconfigure %s -disabledimage %v", el.GetParent().GetID(), el.GetID(), img.GetID())
}