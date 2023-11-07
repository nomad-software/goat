// Code generated by tooling; DO NOT EDIT.
package spinbox

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
)




// SetFont sets the widget's font.
func (el *Spinbox) SetFont(font string, size string, styles ...string) {
	style := strings.Join(styles, " ")
	tk.Get().Eval("%s configure -font {{%s} %s %s}", el.GetID(), font, size, style)
}

// SetFontFromDialog sets the widget's font from the output of the font dialog.
func (el *Spinbox) SetFontFromDialog(dialogOutput string) {
	tk.Get().Eval("%s configure -font {%s}", el.GetID(), dialogOutput)
}
