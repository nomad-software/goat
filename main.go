package main

import (
	"log/slog"

	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/tk"
)

func main() {
	log.SetLevel(slog.LevelDebug)

	tk := tk.Get()
	tk.Eval("wm geometry . 1024x768")
	tk.Start()
}
