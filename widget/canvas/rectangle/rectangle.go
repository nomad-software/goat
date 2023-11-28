package rectangle

import (
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvasrectangle"
)

// Rect represents a rectangle in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/dash
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/fill
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/outline
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/tag
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/width
type Rectangle struct {
	element.Ele
}