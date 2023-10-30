package main

import (
	"github.com/nomad-software/goat/app"
	"github.com/nomad-software/goat/element/color"
	"github.com/nomad-software/goat/example/image"
	"github.com/nomad-software/goat/image/store"
	"github.com/nomad-software/goat/tk/command"
	"github.com/nomad-software/goat/window/protocol"
)

func main() {
	embedded := store.New(image.FS)
	icons := embedded.GetImages("png/tkicon.png")

	app := app.New()
	// app.SetTheme(theme.Clam)
	main := app.GetMainWindow()
	main.SetTitle("Goat showcase")
	main.SetMinSize(600, 600)
	main.SetIcon(icons, true)

	main.BindProtocol(protocol.DeleteWindow, func(*command.CallbackData) {
		//show dialog.
		main.Destroy()
	})

	main.Bind("<Key-Escape>", func(data *command.CallbackData) {
		// fmt.Printf("%#v\n", data)
		app.Exit()
	})

	main.SetBackgroundColor(color.Beige)

	app.Start()
}
