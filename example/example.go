package main

import (
	"github.com/nomad-software/goat/app"
	"github.com/nomad-software/goat/element/compound"
	"github.com/nomad-software/goat/element/relief"
	"github.com/nomad-software/goat/element/underline"
	"github.com/nomad-software/goat/example/image"
	"github.com/nomad-software/goat/image/store"
	"github.com/nomad-software/goat/tk/command"
	"github.com/nomad-software/goat/widget/entry"
	"github.com/nomad-software/goat/widget/frame"
	"github.com/nomad-software/goat/widget/geometry"
	"github.com/nomad-software/goat/widget/labelframe"
	"github.com/nomad-software/goat/widget/menu"
	"github.com/nomad-software/goat/widget/notebook"
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
	createNotebook(main)

	app.Start()
}

func createMenu(win *window.Window) {
	bar := menu.NewBar(win)

	checkSubMenu := menu.NewPopUp()
	checkSubMenu.AddCheckButtonEntry("Option 1", "", func(*command.CallbackData) {})
	checkSubMenu.AddCheckButtonEntry("Option 2", "", func(*command.CallbackData) {})
	checkSubMenu.AddCheckButtonEntry("Option 3", "", func(*command.CallbackData) {})
	checkSubMenu.SetCheckButtonEntry(0, true)

	radioSubMenu := menu.NewPopUp()
	radioSubMenu.AddRadioButtonEntry("Option 1", "", func(*command.CallbackData) {})
	radioSubMenu.AddRadioButtonEntry("Option 2", "", func(*command.CallbackData) {})
	radioSubMenu.AddRadioButtonEntry("Option 3", "", func(*command.CallbackData) {})
	radioSubMenu.SelectRadioButtonEntry(0)

	file := menu.New(bar, "File", underline.None)
	file.AddMenuEntry(checkSubMenu, "Check button submenu", underline.None)
	file.AddMenuEntry(radioSubMenu, "Radio button submenu", underline.None)
	file.AddSeparator()
	img := embedded.GetImage("png/cancel.png")
	file.AddImageEntry(img, compound.Left, "Quit", "Ctrl-Q", func(*command.CallbackData) {
		win.Destroy()
	})

	help := menu.New(bar, "Help", underline.None)
	img = embedded.GetImage("png/help.png")
	help.AddImageEntry(img, compound.Left, "About...", "F1", func(*command.CallbackData) {})
}

func createNotebook(win *window.Window) {
	note := notebook.New(win)
	widgetPane := createWidgetPane()
	panedPane := createPanedPane()
	canvasPane := createCanvasPane()
	dialogPane := createDialogPane()

	img := embedded.GetImage("png/layout_content.png")
	note.AddImageTab(img, compound.Left, "Widgets", underline.None, widgetPane)

	img = embedded.GetImage("png/application_tile_horizontal.png")
	note.AddImageTab(img, compound.Left, "Panes", underline.None, panedPane)

	img = embedded.GetImage("png/shape_ungroup.png")
	note.AddImageTab(img, compound.Left, "Canvas", underline.None, canvasPane)

	img = embedded.GetImage("png/application_double.png")
	note.AddImageTab(img, compound.Left, "Dialogs", underline.None, dialogPane)

	note.Pack(0, 0, geometry.Side.Top, geometry.Fill.Both, geometry.Anchor.Center, true)
}

func createWidgetPane() *frame.Frame {
	frame := frame.New(nil, 0, relief.Flat)

	entryFrame := labelframe.New(frame, "Text entry", underline.None)
	entryFrame.Pack(10, 0, geometry.Side.Top, geometry.Fill.Both, geometry.Anchor.Center, true)

	e := entry.New(entryFrame)
	e.SetValue("lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum ")
	e.Pack(5, 0, geometry.Side.Left, geometry.Fill.Horizontal, geometry.Anchor.NorthWest, true)

	// rangeFrame := labelframe.New(frame, "Progress & Scale", underline.None)
	// rangeFrame.Pack(10, 0, geometry.Side.Bottom, geometry.Fill.Both, geometry.Anchor.Center, true)
	//
	// buttonFrame := labelframe.New(frame, "Buttons", underline.None)
	// buttonFrame.Pack(10, 0, geometry.Side.Left, geometry.Fill.Both, geometry.Anchor.Center, true)
	//
	// checkbuttonFrame := labelframe.New(frame, "Check buttons", underline.None)
	// checkbuttonFrame.Pack(10, 0, geometry.Side.Left, geometry.Fill.Both, geometry.Anchor.Center, true)
	//
	// radiobuttonFrame := labelframe.New(frame, "Radio buttons", underline.None)
	// radiobuttonFrame.Pack(10, 0, geometry.Side.Left, geometry.Fill.Both, geometry.Anchor.Center, true)

	return frame
}

func createPanedPane() *frame.Frame {
	frame := frame.New(nil, 0, relief.Flat)
	return frame
}

func createCanvasPane() *frame.Frame {
	frame := frame.New(nil, 0, relief.Flat)
	return frame
}

func createDialogPane() *frame.Frame {
	frame := frame.New(nil, 0, relief.Flat)
	return frame
}
