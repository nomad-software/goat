package tag

import (
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "texttag"
)

// Tag represents a tag in a text widget.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=text/border
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=text/color
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=text/font
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=text/justify
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=text/margin
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=text/strikethrough
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=text/underline
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
