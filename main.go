package main

import (
	"fmt"

	"github.com/nomad-software/goat/app"
	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/tk/command"
)

func main() {
	app := app.New()
	main := app.GetMainWindow()

	main.SetSize(1024, 768)
	main.SetTitle("test")

	// win2 := window.New(main)
	// win2.SetTitle("test2")
	// win2.Raise(main)
	// win2.Focus(true)

	log.Debug("x pos: %v", main.GetXPos(false))
	log.Debug("x pos: %v", main.GetXPos(true))

	log.Debug("y pos: %v", main.GetYPos(false))
	log.Debug("y pos: %v", main.GetYPos(true))

	main.Bind("<Button-1>", func(pl *command.CallbackPayload) {
		fmt.Printf("%#v\n", pl)
	})

	app.Start()
}
