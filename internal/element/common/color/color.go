package color

import "github.com/nomad-software/goat/tk"

// SetForegroundColor sets the foreground color.
// See [element.color] for color names.
func SetForegroundColor(id string, c string) {
	tk.Get().Eval("%s configure -foreground %s", id, c)
}

// SetBackgroundColor sets the background color.
// See [element.color] for color names.
func SetBackgroundColor(id string, c string) {
	tk.Get().Eval("%s configure -background %s", id, c)
}

// SetInsertColor sets the insert color.
// See [element.color] for color names.
func SetInsertColor(id string, c string) {
	tk.Get().Eval("%s configure -insertbackground %s", id, c)
}
