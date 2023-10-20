package main

import (
	"fmt"
	"log/slog"

	"github.com/nomad-software/goat/element/cursor"
	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/window"
)

func main() {
	log.SetLevel(slog.LevelDebug)

	win := window.GetMain()
	defer win.Show()

	fmt.Printf("main win class: %s\n", win.GetClass())
	win.SetCursor(cursor.Gumby)
	fmt.Printf("win cursor: %s\n", win.GetCursor())

	win.SetSize(1024, 768)
	win.SetTitle("test")

	win2 := window.New(win)
	fmt.Printf("child win class: %s\n", win2.GetClass())
	win2.SetTitle("test2")
}
