package notebook

import (
	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/element"
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/widget"
)

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
func (n *NoteBook) AddImageTab(text string, underline int, img *image.Image, compound string, el element.Element) {
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
func (n *NoteBook) InsertImageTab(index int, text string, underline int, img *image.Image, compound string, el element.Element) {
	tk.Get().Eval("%s insert %d %s -text {%s} -underline %d -image %s -compound {%s}", n.GetID(), index, el.GetID(), text, underline, img.GetID(), compound)
}

// GetNumberOfTabs gets the number of tabs.
func (n *NoteBook) GetNumberOfTabs() int {
	tk.Get().Eval("%s index end", n.GetID())
	return tk.Get().GetIntResult()
}
