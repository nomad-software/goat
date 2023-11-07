package command

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/command"
)

type stub struct{}            // IGNORE
func (el stub) GetID() string { return "." } // IGNORE

// SetCommand set the command to execute on interaction with the widget.
func (el stub) SetCommand(callback command.Callback) {
	name := command.GenerateName(el.GetID())

	tk.Get().CreateCommand(name, callback)
	tk.Get().Eval("%s configure -command %s", el.GetID(), name)
}

// DeleteCommand deletes the command.
func (el stub) DeleteCommand() {
	tk.Get().Eval("%s configure -command {}", el.GetID())

	name := command.GenerateName(el.GetID())
	tk.Get().DeleteCommand(name)
}
