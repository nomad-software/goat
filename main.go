package main

import (
	"github.com/nomad-software/goat/app"
	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/window"
)

func main() {
	app := app.New()
	main := app.GetMainWindow()

	main.SetSize(1024, 768)
	main.SetTitle("test")

	win2 := window.New(main)
	win2.SetTitle("test2")
	win2.Raise(main)
	win2.Focus(true)

	log.Debug("main win cursor y pos: %v", main.GetCursorYPos())

	log.Debug("screen width: %v", main.GetScreenWidth())
	log.Debug("screen height: %v", main.GetScreenHeight())

	log.Debug("x pos: %v", main.GetXPos(false))
	log.Debug("x pos: %v", main.GetXPos(true))

	log.Debug("y pos: %v", main.GetYPos(false))
	log.Debug("y pos: %v", main.GetYPos(true))

	app.Start()
}
