package window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWindow(t *testing.T) {
	win := New(nil)

	assert.Equal(t, "window", win.GetType())
	assert.Equal(t, "Toplevel", win.GetStyle())

	assert.Regexp(t, `^\.window-[A-Z0-9]{1,8}$`, win.GetID())
}

func TestParent(t *testing.T) {
	win := New(nil)
	child := New(win)

	assert.Equal(t, "window", child.GetType())
	assert.Equal(t, "Toplevel", child.GetStyle())

	assert.Regexp(t, `^\.window-[A-Z0-9]{1,8}\.window-[A-Z0-9]{1,8}$`, child.GetID())
}

func TestSize(t *testing.T) {
	win := New(nil)

	win.SetSize(250, 250)
	win.Update()

	assert.Equal(t, 250, win.GetWidth())
	assert.Equal(t, 250, win.GetHeight())
}

func TestGeometry(t *testing.T) {
	win := New(nil)

	win.SetGeometry(350, 350, 150, 150)
	win.Update()

	assert.Equal(t, 350, win.GetWidth())
	assert.Equal(t, 350, win.GetHeight())

	assert.Equal(t, 150, win.GetXPos(false))
	assert.Equal(t, 187, win.GetYPos(false))
}

func TestTitle(t *testing.T) {
	win := New(nil)
	win.SetTitle("foo")

	assert.Equal(t, "foo", win.GetTitle())
}

func TestWaitForVisiblity(t *testing.T) {
	win := New(nil)

	win.SetSize(250, 250)
	win.WaitForVisibility()

	assert.Equal(t, 250, win.GetWidth())
	assert.Equal(t, 250, win.GetHeight())
}

func TestFullScreen(t *testing.T) {
	win := New(nil)
	assert.False(t, win.GetFullScreen())

	win.SetFullScreen(true)
	win.WaitForVisibility()

	assert.True(t, win.GetFullScreen())
}

func TestTopmost(t *testing.T) {
	win := New(nil)
	assert.False(t, win.GetTopmost())

	win.SetTopmost(true)
	win.WaitForVisibility()

	assert.True(t, win.GetTopmost())
}

func TestIconfiy(t *testing.T) {
	win := New(nil)
	win.SetIconify(true)
	win.SetIconify(false)
}

func TestMinMaxSize(t *testing.T) {
	win := New(nil)

	win.SetMinSize(100, 100)
	win.SetMaxSize(200, 200)

	win.SetSize(250, 250)
	win.Update()
	assert.Equal(t, 200, win.GetWidth())
	assert.Equal(t, 200, win.GetHeight())

	win.SetSize(50, 50)
	win.Update()
	assert.Equal(t, 100, win.GetWidth())
	assert.Equal(t, 100, win.GetHeight())
}
