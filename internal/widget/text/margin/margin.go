package margin

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetLeftMargin sets the left margin.
func (el stub) SetLeftMargin(margin int) {
	tk.Get().Eval("%s tag configure {%s} -lmargin1 %d", el.GetParent().GetID(), el.GetID(), margin)
}

// SetLeftWrapMargin sets the left margin of any wrapping text.
func (el stub) SetLeftWrapMargin(margin int) {
	tk.Get().Eval("%s tag configure {%s} -lmargin2 %d", el.GetParent().GetID(), el.GetID(), margin)
}

// SetLeftMarginColor sets the background color of the left margin.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetLeftMarginColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -lmargincolor {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// SetRightMargin sets the left margin.
func (el stub) SetRightMargin(margin int) {
	tk.Get().Eval("%s tag configure {%s} -rmargin %d", el.GetParent().GetID(), el.GetID(), margin)
}

// SetRightMarginColor sets the background color of the left margin.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetRightMarginColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -rmargincolor {%s}", el.GetParent().GetID(), el.GetID(), c)
}
