// Code generated by tooling; DO NOT EDIT.
package entry

import (
	"github.com/nomad-software/goat/internal/tk"
)




// SetForegroundColor sets the foreground color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Entry) SetForegroundColor(c string) {
	tk.Get().Eval("%s configure -foreground {%s}", el.GetID(), c)
}




