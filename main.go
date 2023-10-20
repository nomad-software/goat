package main

import (
	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/window"
)

func main() {
	win := window.GetMain()

	win.SetSize(1024, 768)
	win.SetTitle("test")

	win2 := window.New(win)
	win2.SetTitle("test2")

	log.Debug("main win cursor y pos: %v", win.GetCursorYPos())

	log.Debug("screen width: %v", win.GetScreenWidth())
	log.Debug("screen height: %v", win.GetScreenHeight())

	log.Debug("x pos: %v", win.GetXPos(false))
	log.Debug("x pos: %v", win.GetXPos(true))

	log.Debug("y pos: %v", win.GetYPos(false))
	log.Debug("y pos: %v", win.GetYPos(true))

	win.Show()
}
