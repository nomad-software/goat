package main

import (
	"github.com/nomad-software/goat/app"
	"github.com/nomad-software/goat/element/color"
)

func main() {
	app := app.New()
	// app.SetTheme(theme.Clam)
	main := app.GetMainWindow()

	main.SetSize(1024, 768)
	main.SetTitle("parent")
	main.SetMinSize(200, 200)
	main.SetMaxSize(1024, 768)
	main.WaitForVisibility()

	main.SetBackgroundColor(color.Aquamarine)

	// log.Debug("screen width: %v", main.GetScreenWidth())
	// log.Debug("screen height: %v", main.GetScreenHeight())
	//
	// log.Debug("x pos: %v", main.GetXPos(false))
	// log.Debug("x pos: %v", main.GetXPos(true))
	//
	// log.Debug("y pos: %v", main.GetYPos(false))
	// log.Debug("y pos: %v", main.GetYPos(true))
	//
	// main.Bind("<Button-1>", func(pl *command.CallbackPayload) {
	// 	fmt.Printf("%#v\n", pl)
	// 	// app.Exit()
	// })

	app.Start()
}
