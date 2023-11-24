package canvas

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/option/relief"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "canvas"
)

// Canvas widgets implement structured graphics. A canvas displays any number
// of items, which may be things like rectangles, circles, lines, and text.
// Items may be manipulated (e.g. moved or re-colored) and commands may be
// associated with items in much the same way that the bind command allows
// commands to be bound to widgets.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/canvas.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=borderwidth
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=color
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=height
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=relief
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=scrollbar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=width
type Canvas struct {
	widget.Widget
}

// New creates a new button.
func New(parent element.Element) *Canvas {
	canvas := &Canvas{}
	canvas.SetParent(parent)
	canvas.SetType(Type)

	tk.Get().Eval("canvas %s", canvas.GetID())

	canvas.SetBorderWidth(1)
	canvas.SetRelief(relief.Sunken)

	return canvas
}

// SetSelectionTolerance sets the selection tolerance.
// Specifies a floating-point value indicating how close the mouse cursor must
// be to an item before it is considered to be “inside” the item. Defaults to
// 1.0.
func (el *Canvas) SetSelectionTolerance(tolerance float64) {
	tk.Get().Eval("%s configure -closeenough %v", el.GetID(), tolerance)
}

// SetConfineScrollRegion sets if the scroll region should be confined.
func (el *Canvas) SetConfineScrollRegion(confine bool) {
	tk.Get().Eval("%s configure -confine %v", el.GetID(), confine)
}

// SetScrollRegion sets the scroll region.
// Specifies a list with four coordinates describing the left, top, right, and
// bottom coordinates of a rectangular region. This region is used for
// scrolling purposes and is considered to be the boundary of the information
// in the canvas.
func (el *Canvas) SetScrollRegion(left, top, right, bottom float64) {
	tk.Get().Eval("%s configure -scrollregion [list %v %v %v %v]", el.GetID(), left, top, right, bottom)
}

// SetScrollStep sets the scroll step which specifies an increment for scrolling.
func (el *Canvas) SetScrollStep(step float64) {
	tk.Get().Eval("%s configure -xscrollincrement %v -yscrollincrement %v", el.GetID(), step, step)
}

// GetPosFromScreenPos gets the canvas position from the screen position.
func (el *Canvas) GetPosFromScreenPos(x, y, grid int) (int, int) {
	tk.Get().Eval("%s canvasx %d %d", el.GetID(), x, grid)
	xPos := tk.Get().GetFloatResult()

	tk.Get().Eval("%s canvasy %d %d", el.GetID(), y, grid)
	yPos := tk.Get().GetFloatResult()

	return int(xPos), int(yPos)
}

// SetScanMark sets the scan mark.
// Records x and y and the canvas's current view; used in conjunction with
// later scan dragto commands. Typically this command is associated with a
// mouse button press in the widget and x and y are the coordinates of the
// mouse.
func (el *Canvas) SetScanMark(x, y int) {
	tk.Get().Eval("%s scan mark %d %d", el.GetID(), x, y)
}

// ScanDragTo computes the difference between its x and y arguments (which are
// typically mouse coordinates) and the x and y arguments to the last scan mark
// command for the widget. It then adjusts the view by gain times the
// difference in coordinates, where gain defaults to 10. This command is
// typically associated with mouse motion events in the widget, to produce the
// effect of dragging the canvas at high speed through its window. The return
// value is an empty string.
func (el *Canvas) ScanDragTo(x, y, gain int) {
	tk.Get().Eval("%s scan dragto %d %d %d", el.GetID(), x, y, gain)
}
