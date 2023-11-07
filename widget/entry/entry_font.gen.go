// Code generated by tooling; DO NOT EDIT.
package entry

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
)




// SetFont sets the widget's font.
func (el *Entry) SetFont(font string, size string, styles ...string) {
	style := strings.Join(styles, " ")
	tk.Get().Eval("%s configure -font {{%s} %s %s}", el.GetID(), font, size, style)
}

// SetFontFromDialog sets the widget's font from the output of the font dialog.
func (el *Entry) SetFontFromDialog(dialogOutput string) {
	tk.Get().Eval("%s configure -font {%s}", el.GetID(), dialogOutput)
}
