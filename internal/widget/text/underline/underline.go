package underline

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetUnderline sets whether the text us underlined.
func (el stub) SetUnderline(underline bool) {
	tk.Get().Eval("%s tag configure {%s} -underline %v", el.GetParent().GetID(), el.GetID(), underline)
}

// SetUnderlineColor sets the underline foreground color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetUnderlineColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -underlinefg {%s}", el.GetParent().GetID(), el.GetID(), c)
}
