package main

import (
	"fmt"

	"github.com/nomad-software/goat/app"
	"github.com/nomad-software/goat/app/theme"
	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/tk/command"
)

func main() {
	app := app.New()
	app.SetTheme(theme.Clam)
	main := app.GetMainWindow()

	main.SetSize(1024, 768)
	main.SetTitle("test")
	main.WaitForVisibility()

	log.Debug("screen width: %v", main.GetScreenWidth())
	log.Debug("screen height: %v", main.GetScreenHeight())

	log.Debug("x pos: %v", main.GetXPos(false))
	log.Debug("x pos: %v", main.GetXPos(true))

	log.Debug("y pos: %v", main.GetYPos(false))
	log.Debug("y pos: %v", main.GetYPos(true))

	main.Bind("<Button-1>", func(pl *command.CallbackPayload) {
		fmt.Printf("%#v\n", pl)
		// app.Exit()
	})

	app.Start()
}
