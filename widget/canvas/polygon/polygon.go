package polygon

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvaspolygon"
)

// Polygon represents a polygon in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Polygon -pkg=canvas/dash
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Polygon -pkg=canvas/fill
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Polygon -pkg=canvas/outline
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Polygon -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Polygon -pkg=canvas/tag
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Polygon -pkg=canvas/width
type Polygon struct {
	element.Ele
}

// SetJoinStyle specifies the ways in which joints are to be drawn at the
// vertices of the line. If this option is not specified then it defaults to
// round. If the line only contains two points then this option is irrelevant.
// See [widget.canvas.line.joinstyle] for join style options.
func (el *Polygon) SetJoinStyle(style string) {
	tk.Get().Eval("%s itemconfigure %s -joinstyle {%s}", el.GetParent().GetID(), el.GetID(), style)
}

// SetSmoothMethod sets the smooth method.
// If the smoothing method is bezier, this indicates that the line should be
// drawn as a curve, rendered as a set of quadratic splines: one spline is
// drawn for the first and second line segments, one for the second and third,
// and so on. Straight-line segments can be generated within a curve by
// duplicating the end-points of the desired line segment. If the smoothing
// method is raw, this indicates that the line should also be drawn as a curve
// but where the list of coordinates is such that the first coordinate pair
// (and every third coordinate pair thereafter) is a knot point on a cubic
// Bezier curve, and the other coordinates are control points on the cubic
// Bezier curve. Straight line segments can be generated within a curve by
// making control points equal to their neighbouring knot points. If the last
// point is a control point and not a knot point, the point is repeated (one or
// two times) so that it also becomes a knot point.
// See [widget.canvas.line.smoothmethod] for join style options.
func (el *Polygon) SetSmoothMethod(method string) {
	tk.Get().Eval("%s itemconfigure %s -smooth {%s}", el.GetParent().GetID(), el.GetID(), method)
}

// SetSplineSteps specifies the degree of smoothness desired for curves: each
// spline will be approximated with number line segments. This option is
// ignored unless the smooth method is set.
func (el *Polygon) SetSplineSteps(n int) {
	tk.Get().Eval("%s itemconfigure %s -splinesteps %d", el.GetParent().GetID(), el.GetID(), n)
}
