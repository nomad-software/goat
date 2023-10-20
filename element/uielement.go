package element

import "github.com/nomad-software/goat/tk"

type UIElement interface {
	GetClass() string
	SetCursor(string)
	GetCursor() string
	// Bind(???)
	// UnBind(???)
	Destroy()
	GetWidth() int
	GetHeight() int
	GetCursorPos() []int
	GetXPos(bool) int
	GetYPos(bool) int
	// GenerateEvent(string)
	Focus(bool)
	Lower(Element)
	Raise(Element)
}

// UIElementImpl provides a base implementation of an ui element.
type UIElementImpl struct {
	ElementImpl
}

// GetClass gets the ui element class.
// See [class] package for class names.
func (e *UIElementImpl) GetClass() string {
	tk.Get().Eval("%s cget -class ", e.GetID())

	result := tk.Get().GetResult()

	if result == "" {
		tk.Get().Eval("winfo class %s", e.GetID())
		result = tk.Get().GetResult()
	}

	return result
}

// SetCursor sets the cursor of the ui element.
// See [cursor] package for cursor names.
func (e *UIElementImpl) SetCursor(cursor string) {
	tk.Get().Eval("%s configure -cursor {%s}", e.GetID(), cursor)
}

// GetCursor gets the cursor of the ui element.
// See [cursor] package for cursor names.
func (e *UIElementImpl) GetCursor() string {
	tk.Get().Eval("%s cget -cursor", e.GetID())
	return tk.Get().GetResult()
}

// Destroy removes the element from the UI and cleans up its resources.
// Once destroyed you cannot refer to this ui element again or you will get a
// bad path name error from the interpreter.
func (e *UIElementImpl) Destroy() {
	tk.Get().Eval("destroy %s", e.GetID())
	e.SetType("destroyed")
}
