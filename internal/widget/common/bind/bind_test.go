package bind

import (
	"testing"

	"github.com/nomad-software/goat/internal/tk/command"
	"github.com/stretchr/testify/assert"
)

func TestUIElementBind(t *testing.T) {
	el := stub{}

	el.Bind("<<Modified>>", func(data *command.CallbackData) {
		assert.Equal(t, ".", data.Element.GetID())
	})

	el.GenerateEvent("<<Modified>>")
}

func TestUIElementUnBind(t *testing.T) {
	el := stub{}

	el.Bind("<<Modified>>", func(data *command.CallbackData) {
		assert.Fail(t, "this should have been unbound")
	})

	el.UnBind("<<Modified>>")
	el.GenerateEvent("<<Modified>>")
}
