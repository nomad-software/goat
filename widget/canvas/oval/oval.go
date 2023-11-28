package oval

import (
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvasoval"
)

// Oval represents an oval in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/dash
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/fill
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/outline
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/tag
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/width
type Oval struct {
	element.Ele
}
