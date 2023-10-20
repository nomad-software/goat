package main

import (
	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/window"
)

func main() {
	win := window.GetMain()
	defer win.Show()

	win.SetSize(1024, 768)
	win.SetTitle("test")

	win2 := window.New(win)
	win2.SetTitle("test2")

	log.Debug("child win size: %d, %d", win2.GetWidth(), win2.GetHeight())
}
