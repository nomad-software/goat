package ui

import (
	"testing"

	"github.com/nomad-software/goat/element/cursor"
	"github.com/nomad-software/goat/tk/command"
	"github.com/stretchr/testify/assert"
)

func TestUIElementClass(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.Equal(t, "Tk", el.GetClass())
}

func TestUIElementCursor(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")
	el.SetCursor(cursor.Pirate)

	assert.Equal(t, cursor.Pirate, el.GetCursor())
}

func TestUIElementKeyboardFocus(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.False(t, el.GetKeyboadFocus())

	el.SetKeyboadFocus(true)
	assert.True(t, el.GetKeyboadFocus())
}

func TestUIElementDimensions(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.Equal(t, 1, el.GetWidth())
	assert.Equal(t, 1, el.GetHeight())
}

func TestUIElementOSHandle(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.Greater(t, el.GetOSHandle(), int64(0))
}

func TestUIElementCursorPosition(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	pos := el.GetCursorPos()
	x := el.GetCursorXPos()
	y := el.GetCursorYPos()

	assert.Greater(t, pos[0], 0)
	assert.Greater(t, pos[1], 0)
	assert.Greater(t, x, 0)
	assert.Greater(t, y, 0)

	assert.Equal(t, x, pos[0])
	assert.Equal(t, y, pos[1])
}

func TestUIElementScreenDimensions(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.Greater(t, el.GetScreenWidth(), 0)
	assert.Greater(t, el.GetScreenHeight(), 0)
}

func TestUIElementPosition(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	el.Update()

	assert.Greater(t, el.GetXPos(false), 0)
	assert.Greater(t, el.GetYPos(false), 0)
}

func TestUIElementFocus(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	// Can't really test this here.
	el.Focus(true)
	el.Focus(false)
}

func TestUIElementZPosition(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	// Can't really test this here.
	el.Raise(nil)
	el.Lower(nil)
}

func TestUIElementBind(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	el.Bind("<<Modified>>", func(data *command.CallbackData) {
		assert.Equal(t, ".", data.ElementID)
	})

	el.GenerateEvent("<<Modified>>")
}

func TestUIElementUnBind(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	el.Bind("<<Modified>>", func(data *command.CallbackData) {
		assert.Fail(t, "this should have been unbound")
	})

	el.UnBind("<<Modified>>")
	el.GenerateEvent("<<Modified>>")
}

func TestUIElementDestroy(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.Equal(t, "Tk", el.GetClass())

	el.Destroy()

	assert.Panics(t, func() {
		assert.Equal(t, "Tk", el.GetClass())
	})
}
