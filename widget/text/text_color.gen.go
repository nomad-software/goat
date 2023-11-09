// Code generated by tooling; DO NOT EDIT.
package text

import (
	"github.com/nomad-software/goat/internal/tk"
)




// SetForegroundColor sets the foreground color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Text) SetForegroundColor(c string) {
	tk.Get().Eval("%s configure -foreground {%s}", el.GetID(), c)
}

// SetBackgroundColor sets the background color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Text) SetBackgroundColor(c string) {
	tk.Get().Eval("%s configure -background {%s}", el.GetID(), c)
}

// SetInsertColor sets the insert color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Text) SetInsertColor(c string) {
	tk.Get().Eval("%s configure -insertbackground {%s}", el.GetID(), c)
}
