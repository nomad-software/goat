package command

import (
	"fmt"

	"github.com/nomad-software/goat/element/hash"
)

// Callback is the main command callback that is specified for a command.
type Callback = func(*CallbackPayload)

// CallbackPayload is the main payload of the callback.
// This is automatically loaded with data before the call and is populated
// relevant with data during the call.
type CallbackPayload struct {
	CommandName string
	ElementID   string
	Callback    Callback
	Event       Event
	Dialog      Dialog
}

// Event is the part of the payload that contains information about any events
// that have taken place.
type Event struct {
	MouseButton int
	KeyCode     int
	X           int
	Y           int
	Wheel       int
	Key         string
	ScreenX     int
	ScreenY     int
}

// Dialog is the part of the payload that contain information about dialog
// interaction.
type Dialog struct {
	Font string
}

// GenerateName generates a custom command name.
func GenerateName(args ...string) string {
	args = append(args, "command")
	hash := hash.Generate(args...)

	return fmt.Sprintf("command-%s", hash)
}
