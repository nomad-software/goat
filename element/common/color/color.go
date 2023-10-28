package color

import "github.com/nomad-software/goat/tk"

// SetForegroundColor sets the foreground color.
// See [element.color] for color names.
func SetForegroundColor(id string, color string) {
	tk.Get().Eval("%s configure -foreground %s", id, color)
}

// SetBackgroundColor sets the background color.
// See [element.color] for color names.
func SetBackgroundColor(id string, color string) {
	tk.Get().Eval("%s configure -background %s", id, color)
}

// SetInsertColor sets the insert color.
// See [element.color] for color names.
func SetInsertColor(id string, color string) {
	tk.Get().Eval("%s configure -insertbackground %s", id, color)
}
