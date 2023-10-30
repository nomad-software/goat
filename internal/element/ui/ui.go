package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nomad-software/goat/internal/element"
	"github.com/nomad-software/goat/log"
	"github.com/nomad-software/goat/tk"
	"github.com/nomad-software/goat/tk/command"
)

// Ele provides a base implementation of an ui element.
type Ele struct {
	element.Ele
}

// Update causes operations that are normally deferred, such as display updates
// and window layout calculations, to be performed immediately.
func (e *Ele) Update() {
	tk.Get().Eval("update idletasks")
}

// GetClass gets the ui element class.
// See [element.class] for class names.
func (e *Ele) GetClass() string {
	tk.Get().Eval("%s cget -class ", e.GetID())

	result := tk.Get().GetStrResult()

	if result == "" {
		tk.Get().Eval("winfo class %s", e.GetID())
		result = tk.Get().GetStrResult()
	}

	return result
}

// GetStyle gets the ui element class.
// See [element.style] for style names.
func (e *Ele) GetStyle() string {
	tk.Get().Eval("%s cget -style ", e.GetID())
	return tk.Get().GetStrResult()
}

// SetCursor sets the cursor of the ui element.
// See [element.cursor] for cursor names.
func (e *Ele) SetCursor(cursor string) {
	tk.Get().Eval("%s configure -cursor {%s}", e.GetID(), cursor)
}

// GetCursor gets the cursor of the ui element.
// See [element.cursor] for cursor names.
func (e *Ele) GetCursor() string {
	tk.Get().Eval("%s cget -cursor", e.GetID())
	return tk.Get().GetStrResult()
}

// SetKeyboadFocus sets that this ui element accepts the focus during keyboard
// traversal.
func (e *Ele) SetKeyboadFocus(focus bool) {
	tk.Get().Eval("%s configure -takefocus %v", e.GetID(), focus)
}

// GetKeyboadFocus returns true if this ui element accepts the focus during
// keyboard traversal.
func (e *Ele) GetKeyboadFocus() bool {
	tk.Get().Eval("%s cget -takefocus", e.GetID())
	return tk.Get().GetBoolResult()
}

// Destroy removes the ui element from the UI and cleans up its resources. Once
// destroyed you cannot refer to this ui element again or you will get a bad
// path name error from the interpreter.
func (e *Ele) Destroy() {
	tk.Get().Eval("destroy %s", e.GetID())
	e.SetType("destroyed")
}

// GetWidth gets the width of the ui element.
//
// Returns an int giving a ui element width in pixels. When a ui element is
// first created its width will be 1 pixel; the width will eventually be
// changed by a geometry manager to fulfil the window's needs.
func (e *Ele) GetWidth() int {
	tk.Get().Eval("winfo width %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetHeight gets the height of the ui element.
//
// Returns an int giving a ui element height in pixels. When a ui element
// is first created its height will be 1 pixel; the height will eventually be
// changed by a geometry manager to fulfil the window's needs.
func (e *Ele) GetHeight() int {
	tk.Get().Eval("winfo height %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetOSHandle gets the OS specific window handle.
//
// Returns a low-level platform-specific identifier for a window. On Unix
// platforms, this is the X window identifier. Under Windows, this is the
// Windows HWND. On the Macintosh the value has no meaning outside Tk.
func (e *Ele) GetOSHandle() int64 {
	tk.Get().Eval("winfo id %s", e.GetID())
	result := tk.Get().GetStrResult()

	// Remove the 0x prefix.
	if len(result) > 2 {
		result = result[2:]
	}

	hwnd, err := strconv.ParseInt(result, 16, 0)
	if err != nil {
		log.Error(err)
	}

	return hwnd
}

// GetCursorPos gets the x and y position of the cursor on the ui element.
//
// If the mouse pointer is on the same screen as the ui element, it returns a
// list with two integers, which are the pointer's x and y coordinates measured
// in pixels in the screen's root window. If a virtual root window is in use on
// the screen, the position is computed in the virtual root. If the mouse
// pointer is not on the same screen as ui element then both of the returned
// coordinates are -1.
func (e *Ele) GetCursorPos() []int {
	tk.Get().Eval("winfo pointerxy %s", e.GetID())
	result := tk.Get().GetStrResult()

	strs := strings.Split(result, " ")
	pos := make([]int, 0)

	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Error(err)
		}
		pos = append(pos, i)
	}

	return pos
}

// GetCursorXPos gets the x position of the cursor on the ui element.
//
// If the mouse pointer is on the same screen as the ui element, it returns the
// pointer's x coordinate, measured in pixels in the screen's root window. If a
// virtual root window is in use on the screen, the position is measured in the
// virtual root. If the mouse pointer is not on the same screen as ui element
// then -1 is returned.
func (e *Ele) GetCursorXPos() int {
	tk.Get().Eval("winfo pointerx %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetCursorYPos gets the y position of the cursor on the ui element.
//
// If the mouse pointer is on the same screen as the ui element, it returns the
// pointer's y coordinate, measured in pixels in the screen's root window. If a
// virtual root window is in use on the screen, the position is measured in the
// virtual root. If the mouse pointer is not on the same screen as ui element
// then -1 is returned.
func (e *Ele) GetCursorYPos() int {
	tk.Get().Eval("winfo pointery %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetScreenWidth gets the width of the screen this ui element is on.
func (e *Ele) GetScreenWidth() int {
	tk.Get().Eval("winfo screenwidth %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetScreenHeight gets the height of the screen this ui element is on.
func (e *Ele) GetScreenHeight() int {
	tk.Get().Eval("winfo screenheight %s", e.GetID())
	return tk.Get().GetIntResult()
}

// GetXPos gets the x position of the ui element.
// You may need to wait until the ui element has been updated for this to
// return the correct value.
func (e *Ele) GetXPos(relativeToParent bool) int {
	if relativeToParent {
		tk.Get().Eval("winfo x %s", e.GetID())
	} else {
		tk.Get().Eval("winfo rootx %s", e.GetID())
	}
	return tk.Get().GetIntResult()
}

// GetYPos gets the y position of the ui element.
// You may need to wait until the ui element has been updated for this to
// return the correct value.
func (e *Ele) GetYPos(relativeToParent bool) int {
	if relativeToParent {
		tk.Get().Eval("winfo y %s", e.GetID())
	} else {
		tk.Get().Eval("winfo rooty %s", e.GetID())
	}
	return tk.Get().GetIntResult()
}

// Focus gives focus to the ui element.
func (e *Ele) Focus(force bool) {
	if force {
		tk.Get().Eval("focus -force %s", e.GetID())
	} else {
		tk.Get().Eval("focus %s", e.GetID())
	}
}

// Lower lowers a ui element below another if specified or below all of its
// siblings in the stacking order
func (e *Ele) Lower(el element.Element) {
	if el != nil {
		tk.Get().Eval("lower %s %s", e.GetID(), el.GetID())
	} else {
		tk.Get().Eval("lower %s", e.GetID())
	}
}

// Raise raises a ui element above another if specified or above all of its
// siblings in the stacking order.
func (e *Ele) Raise(el element.Element) {
	if el != nil {
		tk.Get().Eval("raise %s %s", e.GetID(), el.GetID())
	} else {
		tk.Get().Eval("raise %s", e.GetID())
	}
}

// Bind binds a callback to a specific binding.
// Once the callback is called, the argument contains information about the
// event and data from the ui element.
//
// # Bindings
//
// The binding argument specifies a sequence of one or more event patterns,
// with optional white space between the patterns. Each event pattern may take
// one of three forms. In the simplest case it is a single printing ASCII
// character, such as 'a' or '['. The character may not be a space character or
// the character '<'. This form of pattern matches a KeyPress event for the
// particular character. The second form of pattern is longer but more general.
// It has the following syntax.
//
//	<modifier-modifier-type-detail>
//
// The entire event pattern is surrounded by angle brackets. Inside the angle
// brackets are zero or more modifiers, an event type, and an extra piece of
// information (detail) identifying a particular button or keysym. Any of the
// fields may be omitted, as long as at least one of type and detail is
// present. The fields must be separated by white space or dashes (dashes are
// prefered). The third form of pattern is used to specify a user-defined,
// named virtual event. It has the following syntax.
//
//	<<name>>
//
// The entire virtual event pattern is surrounded by double angle brackets.
// Inside the angle brackets is the user-defined name of the virtual event.
// Modifiers, such as Shift or Control, may not be combined with a virtual
// event to modify it. Bindings on a virtual event may be created before the
// virtual event is defined, and if the definition of a virtual event changes
// dynamically, all windows bound to that virtual event will respond
// immediately to the new definition. Some widgets (e.g. menu and text) issue
// virtual events when their internal state is updated in some ways. Please see
// the documentation for each widget for details.
//
// # Modifiers
//
//	Control       Button1, B1      Mod1, M1, Command      Meta, M
//	Alt           Button2, B2      Mod2, M2, Option       Double
//	Shift         Button3, B3      Mod3, M3               Triple
//	Lock          Button4, B4      Mod4, M4               Quadruple
//	Extended      Button5, B5      Mod5, M5
//
// Where more than one value is listed, separated by commas, the values are
// equivalent. Most of the modifiers have the obvious X meanings. For example,
// Button1 requires that button 1 be depressed when the event occurs. For a
// binding to match a given event, the modifiers in the event must include all
// of those specified in the event pattern. An event may also contain
// additional modifiers not specified in the binding. For example, if button 1
// is pressed while the shift and control keys are down, the pattern
// <Control-Button-1> will match the event, but <Mod1-Button-1> will not. If no
// modifiers are specified, then any combination of modifiers may be present in
// the event.
//
// Meta and M refer to whichever of the M1 through M5 modifiers is associated
// with the Meta key(s) on the keyboard (keysyms Meta_R and Meta_L). If there
// are no Meta keys, or if they are not associated with any modifiers, then
// Meta and M will not match any events. Similarly, the Alt modifier refers to
// whichever modifier is associated with the alt key(s) on the keyboard
// (keysyms Alt_L and Alt_R).
//
// The Double, Triple and Quadruple modifiers are a convenience for specifying
// double mouse clicks and other repeated events. They cause a particular event
// pattern to be repeated 2, 3 or 4 times, and also place a time and space
// requirement on the sequence: for a sequence of events to match a Double,
// Triple or Quadruple pattern, all of the events must occur close together in
// time and without substantial mouse motion in between. For example,
// <Double-Button-1> is equivalent to <Button-1><Button-1> with the extra time
// and space requirement.
//
// The Command and Option modifiers are equivalents of Mod1 resp. Mod2, they
// correspond to Macintosh-specific modifier keys.
//
// The Extended modifier is, at present, specific to Windows. It appears on
// events that are associated with the keys on the “extended keyboard”. On a US
// keyboard, the extended keys include the Alt and Control keys at the right of
// the keyboard, the cursor keys in the cluster to the left of the numeric pad,
// the NumLock key, the Break key, the PrintScreen key, and the / and Enter
// keys in the numeric keypad.
//
// # Types
//
// The type field may be any of the standard X event types, with a few extra
// abbreviations. The type field will also accept a couple non-standard X event
// types that were added to better support the Macintosh and Windows platforms.
// Below is a list of all the valid types; where two names appear together,
// they are synonyms.
//
//	Activate                 Destroy            Map
//	ButtonPress, Button      Enter              MapRequest
//	ButtonRelease            Expose             Motion
//	Circulate                FocusIn            MouseWheel
//	CirculateRequest         FocusOut           Property
//	Colormap                 Gravity            Reparent
//	Configure                KeyPress, Key      ResizeRequest
//	ConfigureRequest         KeyRelease         Unmap
//	Create                   Leave              Visibility
//	Deactivate
//
// Most of the above events have the same fields and behaviors as events in the
// X Windowing system. You can find more detailed descriptions of these events
// in any X window programming book. A couple of the events are extensions to
// the X event system to support features unique to the Macintosh and Windows
// platforms.
//
// # Details
//
// The last part of a long event specification is detail. In the case of a
// ButtonPress or ButtonRelease event, it is the number of a button (1-5). If a
// button number is given, then only an event on that particular button will
// match; if no button number is given, then an event on any button will match.
// Note: giving a specific button number is different than specifying a button
// modifier; in the first case, it refers to a button being pressed or
// released, while in the second it refers to some other button that is already
// depressed when the matching event occurs. If a button number is given then
// type may be omitted: if will default to ButtonPress. For example, the
// specifier <1> is equivalent to <ButtonPress-1>.
//
// If the event type is KeyPress or KeyRelease, then detail may be specified in
// the form of an X keysym. Keysyms are textual specifications for particular
// keys on the keyboard; they include all the alphanumeric ASCII characters
// (e.g. “a” is the keysym for the ASCII character “a”), plus descriptions for
// non-alphanumeric characters (“comma”is the keysym for the comma character),
// plus descriptions for all the non-ASCII keys on the keyboard (e.g. “Shift_L”
// is the keysym for the left shift key, and “F1” is the keysym for the F1
// function key, if it exists). The complete list of keysyms is not presented
// here; it is available in other X documentation and may vary from system to
// system. If necessary, you can use the %K notation described below to print
// out the keysym name for a particular key. If a keysym detail is given, then
// the type field may be omitted; it will default to KeyPress. For example,
// <Control-comma> is equivalent to <Control-KeyPress-comma>.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/bind.html
func (e *Ele) Bind(binding string, callback command.Callback) {
	if ok := tk.Binding.MatchString(binding); !ok {
		log.Error(fmt.Errorf("invalid binding: %s", binding))
		return
	}

	name := command.GenerateName(binding, e.GetID())

	tk.Get().CreateCommand(name, callback)
	tk.Get().Eval("bind %s {%s} {%s %%W %%b %%k %%x %%y %%D %%K %%X %%Y}", e.GetID(), binding, name)
}

// UnBind unbinds a command from the passed binding.
func (e *Ele) UnBind(binding string) {
	if ok := tk.Binding.MatchString(binding); !ok {
		log.Error(fmt.Errorf("invalid binding: %s", binding))
		return
	}

	name := command.GenerateName(binding, e.GetID())

	tk.Get().Eval("bind %s {%s} {}", e.GetID(), binding)
	tk.Get().DeleteCommand(name)
}

// GenerateEvent generates the passed event on the ui element.
func (e *Ele) GenerateEvent(event string) {
	if ok := tk.Event.MatchString(event); !ok {
		log.Error(fmt.Errorf("invalid event: %s", event))
		return
	}

	tk.Get().Eval("event generate %s {%s}", e.GetID(), event)
}
