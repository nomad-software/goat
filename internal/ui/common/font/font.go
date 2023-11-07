package font

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetFont sets the widget's font.
func (el stub) SetFont(font string, size string, styles ...string) {
	style := strings.Join(styles, " ")
	tk.Get().Eval("%s configure -font {{%s} %s %s}", el.GetID(), font, size, style)
}

// SetFontFromDialog sets the widget's font from the output of the font dialog.
func (el stub) SetFontFromDialog(dialogOutput string) {
	tk.Get().Eval("%s configure -font {%s}", el.GetID(), dialogOutput)
}
