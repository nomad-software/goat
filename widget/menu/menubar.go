package menu

import (
	"github.com/nomad-software/goat/internal/element"
	"github.com/nomad-software/goat/internal/element/ui"
	"github.com/nomad-software/goat/tk"
)

// Menubar is the bar across the top of a window holding the menu items.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/menu.html
type MenuBar struct {
	ui.Ele
}

// NewBar creates a new menu bar to hold the menu.
// The parent will usually be a window.
func NewBar(parent element.Element) *MenuBar {
	bar := &MenuBar{}
	bar.SetParent(parent)
	bar.SetType("menubar")

	tk.Get().Eval("menu %s -tearoff 0", bar.GetID())
	tk.Get().Eval("%s configure -menu %s", bar.GetParent().GetID(), bar.GetID())

	return bar
}

// DisableMenu disables the menu at the specified index.
func (m *MenuBar) DisableMenu(index int) {
	tk.Get().Eval("%s entryconfigure %d -state disable", m.GetID(), index)
}

// EnableMenu enables the menu at the specified index.
func (m *MenuBar) EnableMenu(index int) {
	tk.Get().Eval("%s entryconfigure %d -state normal", m.GetID(), index)
}
