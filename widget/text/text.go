package text

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/ui/element"
	"github.com/nomad-software/goat/option/wrapmode"
	"github.com/nomad-software/goat/widget"
)

type Text struct {
	widget.Widget
}

// New creates a new text widget. A text widget displays one or more lines of
// text and allows that text to be edited.
//
// Virtual events that can also be bound to.
// <<Modified>>
// <<Selection>>
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=height
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=padding
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=scrollbar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=width
func New(parent element.Element) *Text {
	text := &Text{}
	text.SetParent(parent)
	text.SetType("text")

	tk.Get().Eval("text %s -highlightthickness 0", text.GetID())

	text.EnableUndo(true)
	text.SetUndoLevels(100)
	text.SetWrapMode(wrapmode.Word)

	return text
}

// EnableUndo enables undo functionality.
func (el *Text) EnableUndo(enable bool) {
	tk.Get().Eval("%s configure -undo %v", el.GetID(), enable)
}

// SetUndoLevels sets the maximum amount of undo levels.
func (el *Text) SetUndoLevels(levels int) {
	tk.Get().Eval("%s configure -maxundo %d", el.GetID(), levels)
}

// SetWrapMode sets the text wrap mode.
func (el *Text) SetWrapMode(mode string) {
	tk.Get().Eval("%s configure -wrap {%s}", el.GetID(), mode)
}

// AppendText appends text to the end.
func (el *Text) AppendText(text string) {
	tk.Get().Eval("%s insert end {%s}", el.GetID(), text)
}

// InsertText inserts text at the specified line and character.
func (el *Text) InsertText(line, char int, text string) {
	tk.Get().Eval("%s insert %d.%d {%s}", el.GetID(), line, char, text)
}

// GetText gets the current text.
func (el *Text) GetText() string {
	tk.Get().Eval("%s get 0.0 end", el.GetID())
	return tk.Get().GetStrResult()
}

// Clear clears all the text.
func (el *Text) Clear() {
	tk.Get().Eval("%s delete 0.0 end", el.GetID())
}

// SetText sets the text.
func (el *Text) SetText(text string) {
	el.Clear()
	el.AppendText(text)
}

// Undo undo's the last edit.
func (el *Text) Undo() {
	tk.Get().Eval("%s edit undo", el.GetID())
}

// Redo redo's the last edit.
func (el *Text) Redo() {
	tk.Get().Eval("%s edit redo", el.GetID())
}

// ResetUndo clears all undo's.
func (el *Text) ResetUndo() {
	tk.Get().Eval("%s edit reset", el.GetID())
}

// Cut cuts text to the clipboard.
func (el *Text) Cut() {
	tk.Get().Eval("tk_textCut %s", el.GetID())
}

// Copy copies text to the clipboard.
func (el *Text) Copy() {
	tk.Get().Eval("tk_textCopy %s", el.GetID())
}

// Paste pastes text from the clipboard.
func (el *Text) Paste() {
	tk.Get().Eval("tk_textPaste %s", el.GetID())
}

// See scroll the context to show the specified line and character.
func (el *Text) See(line, char int) {
	tk.Get().Eval("%s see %d.%d", el.GetID(), line, char)
}
