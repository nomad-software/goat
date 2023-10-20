package window

import "github.com/nomad-software/goat/tk"

// MainWindow is the struct representing the main window.
type MainWindow struct {
	Window
}

// GetMain gets the main window of the application.
func GetMain() *MainWindow {
	win := &MainWindow{}
	win.SetID(".")
	win.SetType("window")

	return win
}

// Show shows the main window.
// You call this to start the main application.
func (w *MainWindow) Show() {
	tk.Get().Start()
}
