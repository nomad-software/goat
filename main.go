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

	log.Debug("main win size: %d, %d", win.GetWidth(), win.GetHeight())
	log.Debug("child win size: %d, %d", win2.GetWidth(), win2.GetHeight())
	log.Debug("main win handle: %d", win.GetOSHandle())

	win.Show()
}
