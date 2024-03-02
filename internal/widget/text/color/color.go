package color

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetForegroundColor sets the foreground color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetForegroundColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -foreground {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// SetBackgroundColor sets the background color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetBackgroundColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -background {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// SetSelectForegroundColor sets the selection foreground color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetSelectForegroundColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -selectforeground {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// SetSelectBackgroundColor sets the selection background color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetSelectBackgroundColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -selectbackground {%s}", el.GetParent().GetID(), el.GetID(), c)
}
