package notebook

import (
	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/element"
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/widget"
)

// A notebook widget manages a collection of panes and displays a single one at
// a time. Each pane is associated with a tab, which the user may select to
// change the currently displayed pane.
//
// Virtual events that can also be bound to.
//
// <<NotebookTabChanged>>
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_notebook.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*NoteBook -pkg=height
//go:generate go run ../../internal/tools/generate/main.go -recv=*NoteBook -pkg=padding
//go:generate go run ../../internal/tools/generate/main.go -recv=*NoteBook -pkg=width
type NoteBook struct {
	widget.Widget
}

func New(parent element.Element) *NoteBook {
	note := &NoteBook{}
	note.SetParent(parent)
	note.SetType("notebook")

	tk.Get().Eval("ttk::notebook %s", note.GetID())

	return note
}

// AddTab adds a tab to the end.
// See [element.underline] for underline options.
func (n *NoteBook) AddTab(text string, underline int, el element.Element) {
	tk.Get().Eval("%s insert end %s -text {%s} -underline %d", n.GetID(), el.GetID(), text, underline)
}

// AddImageTab is the same as AddTab but also displays an image.
// See [element.underline] for underline options.
// See [element.compound] for image positions.
func (n *NoteBook) AddImageTab(img *image.Image, compound string, text string, underline int, el element.Element) {
	tk.Get().Eval("%s insert end %s -text {%s} -underline %d -image %s -compound {%s}", n.GetID(), el.GetID(), text, underline, img.GetID(), compound)
}

// InsertTab inserts a tab at the specified index.
// See [element.underline] for underline options.
func (n *NoteBook) InsertTab(index int, text string, underline int, el element.Element) {
	tk.Get().Eval("%s insert %d %s -text {%s} -underline %d", n.GetID(), index, el.GetID(), text, underline)
}

// InsertImageTab is the same as InsertTab but also displays an image.
// See [element.underline] for underline options.
// See [element.compound] for image positions.
func (n *NoteBook) InsertImageTab(img *image.Image, compound string, index int, text string, underline int, el element.Element) {
	tk.Get().Eval("%s insert %d %s -text {%s} -underline %d -image %s -compound {%s}", n.GetID(), index, el.GetID(), text, underline, img.GetID(), compound)
}

// SelectTab selects the tab specified by the passed index.
func (n *NoteBook) SelectTab(index int) {
	count := n.GetNumberOfTabs()
	if index >= count {
		index = count - 1
	}
	tk.Get().Eval("%s select %d", n.GetID(), index)
}

// RemoveTab removes the tab specified by the passed index.
func (n *NoteBook) RemoveTab(index int) {
	count := n.GetNumberOfTabs()
	if index >= count {
		index = count - 1
	}
	tk.Get().Eval("%s forget %d", n.GetID(), index)
}

// DisableTab disables the tab specified by the passed index.
func (n *NoteBook) DisableTab(index int) {
	count := n.GetNumberOfTabs()
	if index >= count {
		index = count - 1
	}
	tk.Get().Eval("%s tab %d -state disable", n.GetID(), index)
}

// EnableTab enables the tab specified by the passed index.
func (n *NoteBook) EnableTab(index int) {
	count := n.GetNumberOfTabs()
	if index >= count {
		index = count - 1
	}
	tk.Get().Eval("%s tab %d -state normal", n.GetID(), index)
}

// SetPaneStickyState sets a tab pane's sticky state. Specifies how the widget
// is positioned within the pane area. Sticky state is a string containing zero
// or more of the characters n, s, e, or w. Each letter refers to a side
// (north, south, east, or west) that the widget will "stick" to, as per the
// grid geometry manager.
func (n *NoteBook) SetPaneStickyState(index int, state string) {
	count := n.GetNumberOfTabs()
	if index >= count {
		index = count - 1
	}
	tk.Get().Eval("%s tab %d -sticky {%s}", n.GetID(), index, state)
}

// SetPanePadding sets the pane's padding.
func (n *NoteBook) SetPanePadding(index int, padding int) {
	count := n.GetNumberOfTabs()
	if index >= count {
		index = count - 1
	}
	tk.Get().Eval("%s tab %d -padding %d", n.GetID(), index, padding)
}

// SetTabText sets the tab's text.
func (n *NoteBook) SetTabText(index int, text string) {
	count := n.GetNumberOfTabs()
	if index >= count {
		index = count - 1
	}
	tk.Get().Eval("%s tab %d -text {%s}", n.GetID(), index, text)
}

// GetNumberOfTabs gets the number of tabs.
func (n *NoteBook) GetNumberOfTabs() int {
	tk.Get().Eval("%s index end", n.GetID())
	return tk.Get().GetIntResult()
}