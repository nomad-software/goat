package app

import (
	"fmt"

	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/window"
)

// init configures the environment.
func init() {
	tk.Get().Eval("encoding system utf-8")
}

// App is the struct representing the application.
type App struct {
}

// New creates the main window of the application.
func New() *App {
	app := &App{}

	return app
}

func (w *App) GetMainWindow() *window.Window {
	win := &window.Window{}
	win.SetID(".")
	win.SetType("window")

	return win
}

// Start shows the main window and starts the application.
// This method should not be deferred in the main function or else it will
// potentially trap panics in other parts of the program.
func (w *App) Start() {
	tk.Get().Start()
}

// SetTheme sets the theme of the app.
// See [app.theme] for theme names.
func (w *App) SetTheme(theme string) {
	tk.Get().Eval("ttk::style theme use {%s}", theme)
}

// Update is used to bring the application 'up to date' by entering the event
// loop repeatedly until all pending events (including idle callbacks) have
// been processed.
func (w *App) Update() {
	tk.Get().Eval("update")
}

// CreateVirtualEvent associates the virtual event with the binding, so that the
// virtual event will trigger whenever the binding occurs. Virtual events may
// be any string value and binding may have any of the values allowed for the
// binding argument of [element.ui.Bind]. If the virtual event is already
// defined, the new binding adds to the existing bindings for the event.
func (w *App) CreateVirtualEvent(event, binding string) {
	if ok := tk.VirtualEvent.MatchString(event); !ok {
		log.Error(fmt.Errorf("invalid virtual event: %s", event))
		tk.Get().Destroy(1)
	}

	if ok := tk.Binding.MatchString(binding); !ok {
		log.Error(fmt.Errorf("invalid binding: %s", binding))
		tk.Get().Destroy(1)
	}

	tk.Get().Eval("event add {%s} {%s}", event, binding)
}

// DeleteVirtualEvent deletes each of the bindings from those associated with
// the virtual event. Virtual events may be any string value and binding may
// have any of the values allowed for the binding argument of
// [element.ui.Bind]. Any bindings not currently associated with virtual events
// are ignored. If no binding argument is provided, all bindings are removed
// for the virtual event, so that the virtual event will not trigger anymore.
func (w *App) DeleteVirtualEvent(event, binding string) {
	if ok := tk.VirtualEvent.MatchString(event); !ok {
		log.Error(fmt.Errorf("invalid virtual event: %s", event))
		tk.Get().Destroy(1)
	}

	if binding != "" {
		if ok := tk.Binding.MatchString(binding); !ok {
			log.Error(fmt.Errorf("invalid binding: %s", binding))
			tk.Get().Destroy(1)
		}

		tk.Get().Eval("event delete {%s} {%s}", event, binding)

	} else {
		tk.Get().Eval("event delete {%s}", event)
	}
}

// Exit closes the app.
func (w *App) Exit() {
	w.GetMainWindow().Destroy()
}
