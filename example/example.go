package main

import (
	"fmt"

	"github.com/nomad-software/goat/app"
	"github.com/nomad-software/goat/element/color"
	"github.com/nomad-software/goat/example/image"
	"github.com/nomad-software/goat/image/store"
	"github.com/nomad-software/goat/tk/command"
)

func main() {
	app := app.New()
	// app.SetTheme(theme.Clam)

	embedded := store.New(image.FS)

	main := app.GetMainWindow()
	icons := embedded.GetImages("png/tkicon.png")
	main.SetIcon(icons, true)

	main.SetSize(1024, 768)
	main.SetTitle("parent")
	main.SetMinSize(200, 200)
	main.SetMaxSize(1024, 768)
	main.WaitForVisibility()

	main.SetBackgroundColor(color.Beige)

	main.Bind("<Key-Escape>", func(data *command.CallbackData) {
		fmt.Printf("%#v\n", data)
		app.Exit()
	})

	app.Start()
}
