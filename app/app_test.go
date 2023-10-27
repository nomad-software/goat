package app

import (
	"testing"
	"time"

	"github.com/nomad-software/goat/app/theme"
	"github.com/nomad-software/goat/tk/command"
	"github.com/stretchr/testify/assert"
)

func TestAppStart(t *testing.T) {
	app := New()
	app.CreateIdleCallback(time.Millisecond, func(data *command.CallbackData) {
		app.Exit()
	})

	app.Start()
}

func TestAppTheme(t *testing.T) {
	app := New()
	assert.Equal(t, app.GetTheme(), theme.Default)

	app.SetTheme(theme.Clam)
	assert.Equal(t, app.GetTheme(), theme.Clam)
}

func TestAppGetMainWindow(t *testing.T) {
	app := New()
	main := app.GetMainWindow()
	assert.Equal(t, main.GetID(), ".")
	assert.Equal(t, main.GetClass(), "Tk")
}

func TestAppCreateVirtualEvent(t *testing.T) {
	app := New()
	app.CreateVirtualEvent("<<quit-app>>", "<Control-Q>")

	main := app.GetMainWindow()
	main.Bind("<<quit-app>>", func(data *command.CallbackData) {
		app.Exit()
	})

	app.CreateIdleCallback(time.Millisecond, func(data *command.CallbackData) {
		main.GenerateEvent("<Control-Q>")
	})

	app.Start()
}

func TestAppDeleteVirtualEvent(t *testing.T) {
	app := New()
	app.CreateVirtualEvent("<<quit-app>>", "<Control-Q>")
	app.CreateVirtualEvent("<<bad-event>>", "<Control-B>")

	main := app.GetMainWindow()
	main.Bind("<<quit-app>>", func(data *command.CallbackData) {
		app.Exit()
	})
	main.Bind("<<bad-event>>", func(data *command.CallbackData) {
		t.Error("<<bad-event>> was not deleted")
	})

	app.DeleteVirtualEvent("<<bad-event>>", "<Control-B>")

	app.CreateIdleCallback(time.Millisecond, func(data *command.CallbackData) {
		main.GenerateEvent("<Control-B>")
		main.GenerateEvent("<Control-Q>")
	})

	app.Start()
}
