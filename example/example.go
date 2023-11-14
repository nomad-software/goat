package main

import (
	"github.com/nomad-software/goat/app"
	"github.com/nomad-software/goat/example/image"
	"github.com/nomad-software/goat/image/store"
	"github.com/nomad-software/goat/internal/tk/command"
	"github.com/nomad-software/goat/option/anchor"
	"github.com/nomad-software/goat/option/compound"
	"github.com/nomad-software/goat/option/fill"
	"github.com/nomad-software/goat/option/orientation"
	"github.com/nomad-software/goat/option/relief"
	"github.com/nomad-software/goat/option/side"
	"github.com/nomad-software/goat/option/underline"
	"github.com/nomad-software/goat/option/wrapmode"
	"github.com/nomad-software/goat/widget/button"
	"github.com/nomad-software/goat/widget/checkbutton"
	"github.com/nomad-software/goat/widget/combobox"
	"github.com/nomad-software/goat/widget/entry"
	"github.com/nomad-software/goat/widget/frame"
	"github.com/nomad-software/goat/widget/labelframe"
	"github.com/nomad-software/goat/widget/menu"
	"github.com/nomad-software/goat/widget/menubutton"
	"github.com/nomad-software/goat/widget/notebook"
	"github.com/nomad-software/goat/widget/panedwindow"
	"github.com/nomad-software/goat/widget/progressbar"
	"github.com/nomad-software/goat/widget/radiobutton"
	"github.com/nomad-software/goat/widget/scale"
	"github.com/nomad-software/goat/widget/scrollbar"
	"github.com/nomad-software/goat/widget/sizegrip"
	"github.com/nomad-software/goat/widget/spinbox"
	"github.com/nomad-software/goat/widget/text"
	"github.com/nomad-software/goat/widget/treeview"
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

	main := app.GetMainWindow()
	main.SetTitle("Goat showcase")
	main.SetMinSize(600, 600)
	main.SetIcon(icons, true)

	main.Bind("<Control-Key-q>", func(*command.CallbackData) {
		main.Destroy()
	})

	// app.CreateIdleCallback(time.Second, func(data *command.CallbackData) {
	// 	timeEntry.SetValue(time.Now().Format(time.RFC3339))
	// 	app.CreateIdleCallback(time.Second, data.Callback)
	// })

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

	sizegrip := sizegrip.New(win)
	sizegrip.Pack(0, 0, side.Bottom, fill.None, anchor.SouthEast, false)
}

func createWidgetPane() *frame.Frame {
	pane := frame.New(nil, 0, relief.Flat)

	entryFrame := labelframe.New(pane, "Text entry", underline.None)
	entryFrame.Pack(10, 0, side.Top, fill.Both, anchor.Center, true)

	textFrame := frame.New(entryFrame, 0, relief.Flat)
	textFrame.Pack(5, 0, side.Bottom, fill.Both, anchor.Center, true)
	textFrame.SetGridColumnWeight(0, 1)
	textFrame.SetGridRowWeight(0, 1)

	hscroll := scrollbar.NewHorizontal(textFrame)
	hscroll.Grid(0, 1, 0, 0, 1, 1, "esw")

	vscroll := scrollbar.NewVertical(textFrame)
	vscroll.Grid(1, 0, 0, 0, 1, 1, "nes")

	textEntry := text.New(textFrame)
	textEntry.Grid(0, 0, 0, 0, 1, 1, "nesw")
	textEntry.SetWidth(0)
	textEntry.SetHeight(0)
	textEntry.SetText("hello")
	textEntry.SetWrapMode(wrapmode.None)
	textEntry.AttachHorizontalScrollbar(hscroll)
	textEntry.AttachVerticalScrollbar(vscroll)

	hscroll.AttachWidget(textEntry)
	vscroll.AttachWidget(textEntry)

	timeEntry = entry.New(entryFrame)
	timeEntry.Pack(5, 0, side.Left, fill.Horizontal, anchor.NorthWest, true)

	spinEntry := spinbox.New(entryFrame)
	spinEntry.SetData("$foo", "[bar]", "\"baz\"", "{qux}")
	spinEntry.SetWrap(true)
	spinEntry.SetWidth(5)
	spinEntry.Pack(5, 0, side.Left, fill.Horizontal, anchor.North, false)

	comboEntry := combobox.New(entryFrame)
	comboEntry.SetData("Option 1", "Option 2", "Option 3")
	comboEntry.SetValue("Option 1")
	comboEntry.Pack(5, 0, side.Left, fill.Horizontal, anchor.NorthWest, true)

	rangeFrame := labelframe.New(pane, "Progress & Scale", underline.None)
	rangeFrame.Pack(10, 0, side.Bottom, fill.Both, anchor.Center, true)

	progressBar := progressbar.New(rangeFrame, orientation.Horizontal)
	progressBar.SetMaxValue(10)
	progressBar.SetValue(4)
	progressBar.Pack(5, 0, side.Top, fill.Horizontal, anchor.Center, true)

	scale := scale.New(rangeFrame, orientation.Horizontal)
	scale.SetFromValue(10)
	scale.SetToValue(0)
	scale.SetValue(4)
	scale.SetCommand(func(*command.CallbackData) {
		progressBar.SetValue(scale.GetValue())
	})
	scale.Pack(5, 0, side.Top, fill.Horizontal, anchor.Center, true)

	buttonFrame := labelframe.New(pane, "Buttons", underline.None)
	buttonFrame.Pack(10, 0, side.Left, fill.Both, anchor.Center, true)

	button1 := button.New(buttonFrame, "Text button")
	button1.Pack(5, 0, side.Top, fill.None, anchor.Center, false)

	button2 := button.New(buttonFrame, "Image button")
	button2.SetImage(embedded.GetImage("png/disk.png"), compound.Left)
	button2.Pack(5, 0, side.Top, fill.None, anchor.Center, false)

	menu := menu.NewPopUp()
	menu.AddEntry("Option 1", "", func(*command.CallbackData) {})
	menu.AddEntry("Option 2", "", func(*command.CallbackData) {})
	menu.AddEntry("Option 3", "", func(*command.CallbackData) {})
	button3 := menubutton.New(buttonFrame, "Menu button", menu)
	button3.Pack(5, 0, side.Top, fill.None, anchor.Center, false)
	button3.SetImage(embedded.GetImage("png/disk.png"), compound.Left)

	checkbuttonFrame := labelframe.New(pane, "Check buttons", underline.None)
	checkbuttonFrame.Pack(10, 0, side.Left, fill.Both, anchor.Center, true)

	checkbutton1 := checkbutton.New(checkbuttonFrame, "Option 1")
	checkbutton1.Check()
	checkbutton1.Pack(5, 0, side.Top, fill.None, anchor.Center, false)

	checkbutton2 := checkbutton.New(checkbuttonFrame, "Option 2")
	checkbutton2.Pack(5, 0, side.Top, fill.None, anchor.Center, false)

	checkbutton3 := checkbutton.New(checkbuttonFrame, "Option 3")
	checkbutton3.Pack(5, 0, side.Top, fill.None, anchor.Center, false)

	radiobuttonFrame := labelframe.New(pane, "Radio buttons", underline.None)
	radiobuttonFrame.Pack(10, 0, side.Left, fill.Both, anchor.Center, true)

	radiobutton1 := radiobutton.New(radiobuttonFrame, "Option 1")
	radiobutton1.Select()
	radiobutton1.Pack(5, 0, side.Top, fill.None, anchor.Center, false)

	radiobutton2 := radiobutton.New(radiobuttonFrame, "Option 2")
	radiobutton2.Pack(5, 0, side.Top, fill.None, anchor.Center, false)

	radiobutton3 := radiobutton.New(radiobuttonFrame, "Option 3")
	radiobutton3.Pack(5, 0, side.Top, fill.None, anchor.Center, false)

	return pane
}

func createPanedPane() *frame.Frame {
	pane := frame.New(nil, 0, relief.Flat)

	panedWindow := panedwindow.New(pane, orientation.Horizontal)
	panedWindow.Pack(10, 0, side.Top, fill.Both, anchor.Center, true)

	tree := treeview.New(panedWindow)
	tree.SetHeading("Directory listing", anchor.West)
	tree.AddNode(treeview.NewNode("Computer"))
	tree.GetNode(0).AddNode(treeview.NewNode("Documents"))
	tree.GetNode(0).GetNode(0).AddNode(treeview.NewNode("Important notes.txt"))
	tree.GetNode(0).GetNode(0).AddNode(treeview.NewNode("The D Programming Language.pdf"))
	panedWindow.AddPane(tree)
	panedWindow.SetPaneWeight(0, 1)

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
