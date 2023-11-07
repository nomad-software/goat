package main

import (
	"time"

	"github.com/nomad-software/goat/app"
	"github.com/nomad-software/goat/example/image"
	"github.com/nomad-software/goat/image/store"
	"github.com/nomad-software/goat/internal/tk/command"
	"github.com/nomad-software/goat/option/anchor"
	"github.com/nomad-software/goat/option/compound"
	"github.com/nomad-software/goat/option/fill"
	"github.com/nomad-software/goat/option/relief"
	"github.com/nomad-software/goat/option/side"
	"github.com/nomad-software/goat/option/underline"
	"github.com/nomad-software/goat/widget/entry"
	"github.com/nomad-software/goat/widget/frame"
	"github.com/nomad-software/goat/widget/labelframe"
	"github.com/nomad-software/goat/widget/menu"
	"github.com/nomad-software/goat/widget/notebook"
	"github.com/nomad-software/goat/widget/spinbox"
	"github.com/nomad-software/goat/window"
	"github.com/nomad-software/goat/window/protocol"
)

var (
	embedded = store.New(image.FS)

	timeEntry *entry.Entry
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

	app.CreateIdleCallback(time.Second, func(data *command.CallbackData) {
		timeEntry.SetStrValue(time.Now().Format(time.RFC3339))
		app.CreateIdleCallback(time.Second, data.Callback)
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

	note.Pack(0, 0, side.Top, fill.Both, anchor.Center, true)
}

func createWidgetPane() *frame.Frame {
	pane := frame.New(nil, 0, relief.Flat)

	entryFrame := labelframe.New(pane, "Text entry", underline.None)
	entryFrame.Pack(10, 0, side.Top, fill.Both, anchor.Center, true)

	timeEntry = entry.New(entryFrame)
	timeEntry.Pack(5, 0, side.Left, fill.Horizontal, anchor.NorthWest, true)

	spinEntry := spinbox.New(entryFrame)
	spinEntry.SetWidth(5)
	spinEntry.SetWrap(true)
	spinEntry.Pack(5, 0, side.Left, fill.Horizontal, anchor.North, false)

	rangeFrame := labelframe.New(pane, "Progress & Scale", underline.None)
	rangeFrame.Pack(10, 0, side.Bottom, fill.Both, anchor.Center, true)

	buttonFrame := labelframe.New(pane, "Buttons", underline.None)
	buttonFrame.Pack(10, 0, side.Left, fill.Both, anchor.Center, true)

	checkbuttonFrame := labelframe.New(pane, "Check buttons", underline.None)
	checkbuttonFrame.Pack(10, 0, side.Left, fill.Both, anchor.Center, true)

	radiobuttonFrame := labelframe.New(pane, "Radio buttons", underline.None)
	radiobuttonFrame.Pack(10, 0, side.Left, fill.Both, anchor.Center, true)

	return pane
}

func createPanedPane() *frame.Frame {
	pane := frame.New(nil, 0, relief.Flat)
	return pane
}

func createCanvasPane() *frame.Frame {
	pane := frame.New(nil, 0, relief.Flat)
	return pane
}

func createDialogPane() *frame.Frame {
	pane := frame.New(nil, 0, relief.Flat)
	return pane
}
