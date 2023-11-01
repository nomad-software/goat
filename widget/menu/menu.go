package menu

import (
	"fmt"

	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/element"
	"github.com/nomad-software/goat/internal/element/hash"
	"github.com/nomad-software/goat/internal/element/ui"
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/tk/command"
	"github.com/nomad-software/goat/tk/variable"
)

// Menubar is the cascading menu that items are selected from.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/menu.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Menu -pkg=bind
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

	menu.radioButtonVar = variable.GenerateName(label, menu.GetID())

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

// AddMenuEntry adds a cascading menu entry.
func (m *Menu) AddMenuEntry(label string, underline int, menu *Menu) {
	origId := menu.GetID()
	menu.SetParent(m)

	// Update the menu id.
	tk.Get().Eval("%s clone %s", origId, menu.GetID())
	tk.Get().Eval("%s add cascade -label {%s} -underline %d -menu %s", m.GetID(), label, underline, menu.GetID())
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

// AddCheckButtonEntry adds an item to the menu that acts as a check button.
// The shortcut will need to be bound using the Bind method.
func (m *Menu) AddCheckButtonEntry(label string, shortcut string, callback command.Callback) {
	varName := variable.GenerateName(label, m.GetID())
	m.checkButtonVars = append(m.checkButtonVars, varName)

	cmdName := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(cmdName, callback)

	tk.Get().Eval("%s add checkbutton -variable %s -label {%s} -accelerator {%s} -command {%s}", m.GetID(), varName, label, shortcut, cmdName)
}

// AddImageCheckButtonEntry is the same as AddCheckButtonEntry but also
// displays an image.
// The shortcut will need to be bound using the Bind method.
// See [element.compound] for image positions.
func (m *Menu) AddImageCheckButtonEntry(label string, shortcut string, img *image.Image, compound string, callback command.Callback) {
	varName := variable.GenerateName(label, m.GetID())
	m.checkButtonVars = append(m.checkButtonVars, varName)

	cmdName := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(cmdName, callback)

	tk.Get().Eval("%s add checkbutton -variable %s -label {%s} -accelerator {%s} -image %s -compound {%s} -command {%s}", m.GetID(), varName, label, shortcut, img.GetID(), compound, cmdName)
}

// AddRadioButtonEntry adds an item to the menu that acts as a radio button.
// The shortcut will need to be bound using the Bind method.
func (m *Menu) AddRadioButtonEntry(label string, shortcut string, callback command.Callback) {
	name := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(name, callback)

	tk.Get().Eval("%s add radiobutton -variable %s -label {%s} -accelerator {%s} -command {%s}", m.GetID(), m.radioButtonVar, label, shortcut, name)
}

// AddImageRadioButtonEntry is the same as AddRadioButtonEntry but also
// displays an image.
// The shortcut will need to be bound using the Bind method.
// See [element.compound] for image positions.
func (m *Menu) AddImageRadioButtonEntry(label string, shortcut string, img *image.Image, compound string, callback command.Callback) {
	name := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(name, callback)

	tk.Get().Eval("%s add radiobutton -variable %s -label {%s} -accelerator {%s} -image %s -compound {%s} -command {%s}", m.GetID(), m.radioButtonVar, label, shortcut, img.GetID(), compound, name)
}

// GetCheckboxEntrySelected gets if the check box entry at the passed index is
// checked or not. The index only applies to check box entries in the menu not
// any other type of entry. If there are no check box entries in the menu this
// method returns false.
func (m *Menu) GetCheckboxEntrySelected(index int) bool {
	if index >= 0 && index < len(m.checkButtonVars) {
		name := m.checkButtonVars[index]
		return tk.Get().GetVariableBoolValue(name)
	}

	return false
}
