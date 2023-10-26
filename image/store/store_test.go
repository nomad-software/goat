package store

import (
	"testing"

	"github.com/nomad-software/goat/example/image"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	store := New(image.FS)

	img := store.GetImage("png/thumbnail.png")
	assert.Regexp(t, `^\.png-[A-Z0-9]{1,8}$`, img.GetID())
	assert.Equal(t, "png", img.GetType())

	img = store.GetImage("gif/thumbnail.gif")
	assert.Regexp(t, `^\.gif-[A-Z0-9]{1,8}$`, img.GetID())
	assert.Equal(t, "gif", img.GetType())
}
