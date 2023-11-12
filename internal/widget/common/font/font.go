package font

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetFont sets the widget's font.
func (el stub) SetFont(font string, size string, styles ...string) {
	style := strings.Join(styles, " ")
	tk.Get().Eval("%s configure -font {{%s} %s %s}", el.GetID(), font, size, style)
}

// SetFontFromDialog sets the widget's font from the output of the font dialog.
func (el stub) SetFontFromDialog(dialogOutput string) {
	tk.Get().Eval("%s configure -font {%s}", el.GetID(), dialogOutput)
}
