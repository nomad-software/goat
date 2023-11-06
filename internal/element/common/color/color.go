package color

import (
	"github.com/nomad-software/goat/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetForegroundColor sets the foreground color.
// See [widget.option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetForegroundColor(c string) {
	tk.Get().Eval("%s configure -foreground {%s}", el.GetID(), c)
}

// SetBackgroundColor sets the background color.
// See [widget.option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetBackgroundColor(c string) {
	tk.Get().Eval("%s configure -background {%s}", el.GetID(), c)
}

// SetInsertColor sets the insert color.
// See [widget.option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetInsertColor(c string) {
	tk.Get().Eval("%s configure -insertbackground {%s}", el.GetID(), c)
}
