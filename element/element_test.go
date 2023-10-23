package element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverrideID(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.Equal(t, "window", el.GetType())
	assert.Equal(t, ".", el.GetID())
}

func TestGenerateID(t *testing.T) {
	el := &Ele{}
	el.SetType("window")

	assert.Regexp(t, `^\.window-[A-Z0-9]{1,8}$`, el.GetID())
}

func TestMainWindowParent(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	child := &Ele{}
	child.SetType("window")
	child.SetParent(el)

	assert.Implements(t, (*Element)(nil), child.GetParent())
	assert.Regexp(t, `^\.window-[A-Z0-9]{1,8}$`, child.GetID())
}

func TestParent(t *testing.T) {
	el := &Ele{}
	el.SetType("window")

	child := &Ele{}
	child.SetType("window")
	child.SetParent(el)

	assert.Implements(t, (*Element)(nil), child.GetParent())
	assert.Regexp(t, `^\.window-[A-Z0-9]{1,8}\.window-[A-Z0-9]{1,8}$`, child.GetID())
}
