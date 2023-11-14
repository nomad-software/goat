package treeview

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type     = "treeview"
	NodeType = "treeviewnode"
)

// The treeview widget displays a hierarchical collection of items. Each item
// has a textual label, an optional image, and an optional list of data values.
//
// There are two varieties of columns. The first is the main tree view column
// that is present all the time. The second are data columns that can be added
// when needed.
//
// Each tree item has a list of tags, which can be used to associate event
// bindings and control their appearance. Treeview widgets support horizontal
// and vertical scrolling with the standard scroll commands.
//
// Virtual events that can also be bound to.
// <<TreeviewSelect>>
// <<TreeviewOpen>>
// <<TreeviewClose>>
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_treeview.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=height
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=padding
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=scrollbar
type TreeView struct {
	widget.Widget

	nodes []*Node
}

// New creates a new tree view.
func New(parent element.Element) *TreeView {
	tree := &TreeView{}
	tree.SetParent(parent)
	tree.SetType(Type)

	tk.Get().Eval("ttk::treeview %s -selectmode {browse}", tree.GetID())

	return tree
}

// EnableHeading controls showing the heading.
func (el *TreeView) EnableHeading(enable bool) {
	if enable {
		tk.Get().Eval("%s configure -show {tree headings}", el.GetID())
	} else {
		tk.Get().Eval("%s configure -show {tree}", el.GetID())
	}
}

// SetHeading sets the heading.
// See [option.anchor] for anchor values.
func (el *TreeView) SetHeading(text, anchor string) {
	tk.Get().Eval("%s heading #0 -text {%s} -anchor {%s}", el.GetID(), text, anchor)
}

func (el *TreeView) GetNode(index int) *Node {
	return el.nodes[index]
}

func (el *TreeView) AddNode(node *Node) {
	node.SetParent(el)
	el.nodes = append(el.nodes, node)

	tk.Get().Eval("%s insert {} end -text {%s} -values {%s} -open %v -tags [list %s]", el.GetID(), node.GetText(), node.GetText(), true, "foo")

	node.SetID(tk.Get().GetStrResult())
}

// Node represents a node in the tree view.
type Node struct {
	element.Ele

	text  string
	nodes []*Node
}

func NewNode(text string) *Node {
	node := &Node{
		text:  text,
		nodes: make([]*Node, 0),
	}
	node.SetType(NodeType)

	return node
}

func (el *Node) GetText() string {
	return el.text
}

func (el *Node) GetNode(index int) *Node {
	return el.nodes[index]
}

func (el *Node) AddNode(node *Node) {
	node.SetParent(el.GetParent())
	el.nodes = append(el.nodes, node)

	tk.Get().Eval("%s insert %s end -text {%s} -values {%s} -open %v -tags [list %s]", el.GetParent().GetID(), el.GetID(), node.GetText(), node.GetText(), true, "foo")

	node.SetID(tk.Get().GetStrResult())
}
