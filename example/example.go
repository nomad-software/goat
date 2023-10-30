package main

import (
	"fmt"

	"github.com/nomad-software/goat/app"
	"github.com/nomad-software/goat/element/color"
	"github.com/nomad-software/goat/element/compound"
	"github.com/nomad-software/goat/example/image"
	"github.com/nomad-software/goat/image/store"
	"github.com/nomad-software/goat/tk/command"
	"github.com/nomad-software/goat/widget/menu"
	"github.com/nomad-software/goat/window"
	"github.com/nomad-software/goat/window/protocol"
)

var (
	embedded = store.New(image.FS)
)

func main() {
	icons := embedded.GetImages("png/tkicon.png")

	app := app.New()
	// app.SetTheme(theme.Clam)
	main := app.GetMainWindow()
	main.SetTitle("Goat showcase")
	main.SetMinSize(600, 600)
	main.SetIcon(icons, true)

	main.Bind("<Control-Key-q>", func(*command.CallbackData) {
		main.Destroy()
	})

	main.BindProtocol(protocol.DeleteWindow, func(*command.CallbackData) {
		main.Destroy()
	})

	createMenu(main)

	main.SetBackgroundColor(color.Beige)

	app.Start()
}

func createMenu(win *window.Window) {
	bar := menu.NewBar(win)

	checkSubMenu := menu.NewPopUp()
	checkSubMenu.AddCheckboxEntry("Option 1", "", func(*command.CallbackData) {})
	checkSubMenu.AddCheckboxEntry("Option 2", "", func(*command.CallbackData) {})
	checkSubMenu.AddCheckboxEntry("Option 3", "", func(*command.CallbackData) {})

	file := menu.New(bar, "File", 0)
	file.AddMenuEntry("Checkbox submenu", 0, checkSubMenu)

	img := embedded.GetImage("png/cancel.png")
	file.AddImageEntry("Quit", "Ctrl-Q", img, compound.Left, func(*command.CallbackData) {
		win.Destroy()
	})

	fmt.Printf("%v\n", file)
}
