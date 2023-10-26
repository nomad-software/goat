package image

import (
	"github.com/nomad-software/goat/element"
	"github.com/nomad-software/goat/tk"
)

var (
	// Enforce that Img implements Image.
	_ Image = &Img{}
)

// Image defines an image at the lowest level.
type Image interface {
	element.Element
	Blank()
	SetGamma(gamma float64)
	GetGamma() float64
	SetWidth(width int)
	GetWidth() int
	SetHeight(height int)
	GetHeight() int
	Destroy()
}

// Img provides a base implementation of an image.
type Img struct {
	element.Ele
}

// Blank clears the image of all data.
func (i *Img) Blank() {
	tk.Get().Eval("%s blank", i.GetID())
}

// SetGamma sets the gamma.
func (i *Img) SetGamma(gamma float64) {
	tk.Get().Eval("%s configure -gamma %v", i.GetID(), gamma)
}

// GetGamma gets the gamma.
func (i *Img) GetGamma() float64 {
	tk.Get().Eval("%s cget -gamma", i.GetID())
	return tk.Get().GetFloatResult()
}

// SetWidth sets the width.
func (i *Img) SetWidth(width int) {
	tk.Get().Eval("%s configure -width %d", i.GetID(), width)
}

// GetWidth gets the width.
func (i *Img) GetWidth() int {
	tk.Get().Eval("%s cget -width", i.GetID())
	return tk.Get().GetIntResult()
}

// SetHeight sets the height.
func (i *Img) SetHeight(height int) {
	tk.Get().Eval("%s configure -height %d", i.GetID(), height)
}

// GetHeight gets the height.
func (i *Img) GetHeight() int {
	tk.Get().Eval("%s cget -height", i.GetID())
	return tk.Get().GetIntResult()
}

// Destroy deletes the image and cleans up its resources. Once destroyed you
// cannot refer to this image again or you will get a bad path name error from
// the interpreter.
func (i *Img) Destroy() {
	tk.Get().Eval("image delete %s", i.GetID())
}
