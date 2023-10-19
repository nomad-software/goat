package element

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"strings"

	"github.com/nomad-software/goat/tk"
)

// Element defines an element at the lowest level.
type Element interface {
	GetTk() *tk.Tk
	SetHash(string)
	SetID(string)
	GetID() string
	SetParent(Element)
	GetParent() Element
	SetType(id string)
	GenerateHash(args ...string) string
}

// ElementImpl provides a base implementation of the element.
type ElementImpl struct {
	parent       Element
	overriddenID string
	elementID    string
	hash         string
}

// GetTk is a helper method to get the underlying interpreter.
func (e *ElementImpl) GetTk() *tk.Tk {
	return tk.Get()
}

// SetHash sets the element hash.
func (e *ElementImpl) SetHash(hash string) {
	e.hash = hash
}

// GenerateHash generates the element hash.
func (e *ElementImpl) GenerateHash(args ...string) string {
	var text string

	if len(args) > 0 {
		text = strings.Join(args, "")
	} else {
		text = fmt.Sprint(rand.Int63())
	}

	hash := fnv.New32a()
	hash.Write([]byte(text))
	return fmt.Sprintf("%X", hash.Sum32())
}

// SetID sets and overrides the element id.
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

	return parentID + "." + e.elementID + "-" + e.hash
}

// SetParent sets the element parent.
func (e *ElementImpl) SetParent(el Element) {
	e.parent = el
}

// GetParent gets the element parent.
func (e *ElementImpl) GetParent() Element {
	return e.parent
}

// SetType sets the element type. This is used when generating the id.
func (e *ElementImpl) SetType(id string) {
	e.elementID = id
}
