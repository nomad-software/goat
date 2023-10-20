package command

import (
	"fmt"

	"github.com/nomad-software/goat/element/hash"
)

type Callback = func(*CallbackArgs)

type CallbackArgs struct {
	UniqueData string
	Callback   Callback
	Event      Event
	Dialog     Dialog
}

type Event struct {
	Button  int
	KeyCode int
	X       int
	Y       int
	Wheel   int
	ScreenX int
	ScreenY int
}

type Dialog struct {
	font string
}

// GenerateName generates a custom command name.
func GenerateName(args ...string) string {
	args = append(args, "command")
	hash := hash.Generate(args...)

	return fmt.Sprintf("command-%s", hash)
}
