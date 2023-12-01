package tag

import (
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvastag"
)

// Tag represents a tag in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/anchor
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/bind
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/dash
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/fill
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/outline
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/width
type Tag struct {
	element.Ele
}

// Creates a new tag.
func New(parent element.Element) *Tag {
	tag := &Tag{}
	tag.SetParent(parent)
	tag.SetType(Type)

	return tag
}