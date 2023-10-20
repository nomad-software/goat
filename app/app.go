package app

import (
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/window"
)

// init configures the environment.
func init() {
	tk.Get().Eval("encoding system utf-8")
}

// App is the struct representing the application.
type App struct {
}

// New creates the main window of the application.
func New() *App {
	app := &App{}

	return app
}

func (w *App) GetMainWindow() *window.Window {
	win := &window.Window{}
	win.SetID(".")
	win.SetType("window")

	return win
}

// Start shows the main window and starts the application.
// This method should not be deferred in the main function or else it will
// potentially trap panics in other parts of the program.
func (w *App) Start() {
	tk.Get().Start()
}
