package menu

import (
	"fmt"

	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/element"
	"github.com/nomad-software/goat/internal/element/hash"
	"github.com/nomad-software/goat/internal/element/ui"
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/tk/command"
)

// Menubar is the cascading menu that items are selected from.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/menu.html
type Menu struct {
	ui.Ele

	checkButtonVars []string
	radioButtonVar  string
}

// New creates a new menu.
// The parent will usually be a menu bar.
func New(bar element.Element, label string, underline int) *Menu {
	menu := &Menu{}
	menu.SetParent(bar)
	menu.SetType("menu")

	menu.radioButtonVar = fmt.Sprintf("var-%s", hash.Generate())

	tk.Get().Eval("menu %s -type normal -tearoff 0", menu.GetID())
	tk.Get().Eval("%s add cascade -menu %s -label {%s} -underline %d", menu.GetParent().GetID(), menu.GetID(), label, underline)

	return menu
}

// NewPopUp creates a new popup menu that doesn't have a bar as a parent.
func NewPopUp() *Menu {
	menu := &Menu{}
	menu.SetType("popup-menu")

	menu.radioButtonVar = fmt.Sprintf("var-%s", hash.Generate())

	tk.Get().Eval("menu %s -type normal -tearoff 0", menu.GetID())

	return menu
}

// AddEntry adds a menu entry with an optional cosmetic shortcut and a callback
// to execute when selected.
// The shortcut will need to be bound using the Bind method.
func (m *Menu) AddEntry(label string, shortcut string, callback command.Callback) {
	name := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(name, callback)

	tk.Get().Eval("%s add command -label {%s} -accelerator {%s} -command {%s}", m.GetID(), label, shortcut, name)
}

// AddImageEntry is the same as AddEntry but also displays an image.
// The shortcut will need to be bound using the Bind method.
// See [element.compound] for image positions.
func (m *Menu) AddImageEntry(label string, shortcut string, img *image.Image, compound string, callback command.Callback) {
	name := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(name, callback)

	tk.Get().Eval("%s add command -label {%s} -accelerator {%s} -image %s -compound {%s} -command {%s}", m.GetID(), label, shortcut, img.GetID(), compound, name)
}
