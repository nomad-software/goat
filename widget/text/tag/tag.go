package tag

import (
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "texttag"
)

// Tag represents a tag in a text widget.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=text/font
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
