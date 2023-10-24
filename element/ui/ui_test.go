package ui

import (
	"testing"

	"github.com/nomad-software/goat/element/cursor"
	"github.com/nomad-software/goat/tk/command"
	"github.com/stretchr/testify/assert"
)

func TestClass(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.Equal(t, "Tk", el.GetClass())
}

func TestCursor(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")
	el.SetCursor(cursor.Pirate)

	assert.Equal(t, cursor.Pirate, el.GetCursor())
}

func TestKeyboardFocus(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.False(t, el.GetKeyboadFocus())

	el.SetKeyboadFocus(true)
	assert.True(t, el.GetKeyboadFocus())
}

func TestDimensions(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.Equal(t, 1, el.GetWidth())
	assert.Equal(t, 1, el.GetHeight())
}

func TestOSHandle(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.Greater(t, el.GetOSHandle(), int64(0))
}

func TestCursorPosition(t *testing.T) {
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

func TestScreenDimensions(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.Greater(t, el.GetScreenWidth(), 0)
	assert.Greater(t, el.GetScreenHeight(), 0)
}

func TestPosition(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	el.Update()

	assert.Greater(t, el.GetXPos(false), 0)
	assert.Greater(t, el.GetYPos(false), 0)
}

func TestFocus(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	// Can't really test this here.
	el.Focus(true)
	el.Focus(false)
}

func TestZPosition(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	// Can't really test this here.
	el.Raise(nil)
	el.Lower(nil)
}

func TestBind(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	el.Bind("<<Modified>>", func(data *command.CallbackData) {
		assert.Equal(t, ".", data.ElementID)
	})

	el.GenerateEvent("<<Modified>>")
}

func TestUnBind(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	el.Bind("<<Modified>>", func(data *command.CallbackData) {
		assert.Fail(t, "this should have been unbound")
	})

	el.UnBind("<<Modified>>")
	el.GenerateEvent("<<Modified>>")
}

func TestDestroy(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	el.Destroy()

	assert.Panics(t, func() {
		assert.Equal(t, "Tk", el.GetClass())
	})
}
