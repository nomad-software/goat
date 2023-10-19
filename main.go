package main

import (
	"log/slog"

	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/window"
)

func main() {
	log.SetLevel(slog.LevelDebug)

	win := window.GetMain()
	win.SetSize(1024, 768)
	win.SetTitle("test")

	win2 := window.New(win)
	win2.SetTitle("test2")

	win.Show()
}
