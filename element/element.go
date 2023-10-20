package element

import (
	"fmt"

	"github.com/nomad-software/goat/element/hash"
)

// Element defines an element at the lowest level.
type Element interface {
	SetID(string)
	GetID() string
	SetParent(Element)
	GetParent() Element
	SetType(id string)
	GetType() string
}

// ElementImpl provides a base implementation of the element.
type ElementImpl struct {
	parent       Element
	overriddenID string
	typ          string
	hash         string
}

// SetID sets and overrides the element id.
// This should only be called in a constructor.
func (e *ElementImpl) SetID(id string) {
	e.overriddenID = id
}

// GetID gets the element id.
func (e *ElementImpl) GetID() string {
	if e.overriddenID != "" {
		return e.overriddenID
	}

	var parentID string

	if e.parent != nil && e.parent.GetID() != "." {
		parentID = e.parent.GetID()
	}

	if e.hash == "" {
		e.hash = hash.Generate()
	}

	return fmt.Sprintf("%s.%s-%s", parentID, e.typ, e.hash)
}

// SetParent sets the element parent.
// This should only be called in a constructor.
func (e *ElementImpl) SetParent(el Element) {
	e.parent = el
}

// GetParent gets the element parent.
func (e *ElementImpl) GetParent() Element {
	return e.parent
}

// SetType sets the element type.
// This should only be called in a constructor.
func (e *ElementImpl) SetType(id string) {
	e.typ = id
}

// GetType gets the element type.
func (e *ElementImpl) GetType() string {
	return e.typ
}
