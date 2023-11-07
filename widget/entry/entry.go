package entry

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/ui/element"
	"github.com/nomad-software/goat/widget"
)

// An entry widget displays a one-line text string and allows that string to be
// edited by the user. Entry widgets support horizontal scrolling.
//
// Virtual events that can also be bound to.
// <<Clear>>
// <<Copy>>
// <<Cut>>
// <<Paste>>
// <<PasteSelection>>
// <<PrevWindow>>
// <<TraverseIn>>
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_entry.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=color -methods=SetForegroundColor
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=font
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=justify
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=scrollbar -methods=AttachHorizontalScrollbar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=show
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=value
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=width
type Entry struct {
	widget.Widget

	valueVar string
}

// New creates a new entry.
func New(parent element.Element) *Entry {
	entry := &Entry{}
	entry.SetParent(parent)
	entry.SetType("entry")

	entry.valueVar = variable.GenerateName(entry.GetID())

	tk.Get().Eval("ttk::entry %s -textvariable %s", entry.GetID(), entry.valueVar)

	return entry
}
