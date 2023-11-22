package messagedialog

import (
	"github.com/nomad-software/goat/dialog/dialogbutton"
	"github.com/nomad-software/goat/dialog/dialogicon"
	"github.com/nomad-software/goat/dialog/dialogtype"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "messagedialog"
)

// MessageDialog is a dialog box with a user defined message and buttons.
type MessageDialog struct {
	element.Ele

	title         string
	message       string
	detail        string
	dialogType    string
	icon          string
	defaultButton string

	value string
}

// New creates a new message dialog.
func New(parent element.Element, title string) *MessageDialog {
	dialog := &MessageDialog{}

	dialog.SetParent(parent)
	dialog.SetType(Type)

	dialog.SetTitle(title)
	dialog.SetDialogType(dialogtype.Ok)
	dialog.SetIcon(dialogicon.Info)
	dialog.SetDefaultButton(dialogbutton.Ok)

	return dialog
}

// SetTitle sets the dialog title.
func (el *MessageDialog) SetTitle(title string) {
	el.title = title
}

// SetMessage sets the dialog message.
func (el *MessageDialog) SetMessage(msg string) {
	el.message = msg
}

// SetDetail sets the dialog detail.
func (el *MessageDialog) SetDetail(detail string) {
	el.detail = detail
}

// SetDialogType sets the dialog type.
// See [dialog.dialogtype] for dialog type values.
func (el *MessageDialog) SetDialogType(typ string) {
	switch typ {
	case dialogtype.AbortRetryIgnore:
		el.defaultButton = dialogbutton.Abort

	case dialogtype.RetryCancel:
		el.defaultButton = dialogbutton.Retry

	case dialogtype.YesNo:
		el.defaultButton = dialogbutton.Yes

	case dialogtype.YesNoCancel:
		el.defaultButton = dialogbutton.Yes

	default:
		el.defaultButton = dialogbutton.Ok
	}

	el.dialogType = typ
}

// SetIcon sets the dialog icon.
// See [dialog.dialogicon] for icon values.
func (el *MessageDialog) SetIcon(icon string) {
	el.icon = icon
}

// SetDefaultButton sets the dialog default button.
// See [dialog.dialogbutton] for button values.
func (el *MessageDialog) SetDefaultButton(button string) {
	el.defaultButton = button
}

// Show creates and shows the dialog.
func (el *MessageDialog) Show() {
	tk.Get().Eval(
		"tk_messageBox -parent %s -title {%s} -message {%s} -detail {%s} -type {%s} -icon {%s} -default {%s}",
		el.GetParent().GetID(),
		el.title,
		el.message,
		el.detail,
		el.dialogType,
		el.icon,
		el.defaultButton,
	)

	el.value = tk.Get().GetStrResult()
}

// GetValue gets the dialog value when closed.
func (el *MessageDialog) GetValue() string {
	return el.value
}
