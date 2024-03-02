package strikethrough

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetStrikeThrough sets whether the text has a strikethrough.
func (el stub) SetStrikeThrough(underline bool) {
	tk.Get().Eval("%s tag configure {%s} -overstrike %v", el.GetParent().GetID(), el.GetID(), underline)
}

// SetStrikeThroughColor sets the strikethrough foreground color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetStrikeThroughColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -overstrikefg {%s}", el.GetParent().GetID(), el.GetID(), c)
}
