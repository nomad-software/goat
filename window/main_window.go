package window

// MainWindow is the struct representing the main window.
type MainWindow struct {
	Window
}

// GetMain gets the main window of the application.
func GetMain() *MainWindow {
	win := &MainWindow{}
	win.SetHash(win.GenerateHash())
	win.SetType("window")
	win.SetID(".")

	return win
}

// Show shows the main window.
// You call this to start the main application.
func (w *MainWindow) Show() {
	w.GetTk().Start()
}
