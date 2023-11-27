package canvas

import (
	img "github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/option/relief"
	"github.com/nomad-software/goat/widget"
	"github.com/nomad-software/goat/widget/canvas/arc"
	"github.com/nomad-software/goat/widget/canvas/arc/style"
	"github.com/nomad-software/goat/widget/canvas/image"
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
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/borderwidth
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/color
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/height
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/relief
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/scrollbar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/width
type Canvas struct {
	widget.Widget

	itemRef map[string]element.Element
}

// New creates a new button.
func New(parent element.Element) *Canvas {
	canvas := &Canvas{
		itemRef: make(map[string]element.Element),
	}
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

// GetXPosFromScreenXPos gets the x canvas position from the screen position.
func (el *Canvas) GetXPosFromScreenXPos(x, grid int) int {
	tk.Get().Eval("%s canvasx %d %d", el.GetID(), x, grid)
	xPos := tk.Get().GetFloatResult()

	return int(xPos)
}

// GetYPosFromScreenYPos gets the y canvas position from the screen position.
func (el *Canvas) GetYPosFromScreenYPos(y, grid int) int {
	tk.Get().Eval("%s canvasy %d %d", el.GetID(), y, grid)
	yPos := tk.Get().GetFloatResult()

	return int(yPos)
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

// AddArc adds an arc to the canvas.
// The first four coordinates specify the oval that this arc is drawn on.
func (el *Canvas) AddArc(x1, y1, x2, y2 float64) *arc.Arc {
	tk.Get().Eval("%s create arc %v %v %v %v", el.GetID(), x1, y1, x2, y2)
	id := tk.Get().GetStrResult()

	a := &arc.Arc{}
	a.SetParent(el)
	a.SetType(arc.Type)
	a.SetID(id)

	a.SetStyle(style.Pie)

	el.itemRef[id] = a

	return a
}

// AddImage adds an image to the canvas.
func (el *Canvas) AddImage(img *img.Image, x, y float64) *image.Image {
	tk.Get().Eval("%s create image %v %v -image %s", el.GetID(), x, y, img.GetID())
	id := tk.Get().GetStrResult()

	i := &image.Image{}
	i.SetParent(el)
	i.SetType(image.Type)
	i.SetID(id)

	el.itemRef[id] = i

	return i
}
