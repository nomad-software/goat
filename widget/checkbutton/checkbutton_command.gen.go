// Code generated by tooling; DO NOT EDIT.
package checkbutton

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/command"

)



// SetCommand set the command to execute on interaction with the widget.
func (el *CheckButton) SetCommand(callback command.Callback) {
	name := command.GenerateName(el.GetID())

	tk.Get().CreateCommand(el, name, callback)
	tk.Get().Eval("%s configure -command %s", el.GetID(), name)
}

// DeleteCommand deletes the command.
func (el *CheckButton) DeleteCommand() {
	tk.Get().Eval("%s configure -command {}", el.GetID())

	name := command.GenerateName(el.GetID())
	tk.Get().DeleteCommand(name)
}
