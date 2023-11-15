package treeview

import (
	"strings"

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

	reference map[string]*Node
	nodes     []*Node
}

// New creates a new tree view.
// See [option.selectionmode] for mode values.
func New(parent element.Element, selectionMode string) *TreeView {
	tree := &TreeView{
		reference: make(map[string]*Node),
		nodes:     make([]*Node, 0),
	}
	tree.SetParent(parent)
	tree.SetType(Type)

	tk.Get().Eval("ttk::treeview %s -selectmode {%s}", tree.GetID(), selectionMode)

	return tree
}

// SetSelectionMode sets the selection mode of the nodes.
// See [option.selectionmode] for mode values.
func (el *TreeView) SetSelectionMode(mode string) {
	tk.Get().Eval("%s configure -selectmode {%s}", el.GetID(), mode)
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

// AddNode adds a node to the tree view.
func (el *TreeView) AddNode(text, value string, open bool, tags ...string) *Node {
	node := &Node{
		nodes: make([]*Node, 0),
	}

	node.SetParent(el)

	tagStr := strings.Join(tags, " ")
	tk.Get().Eval("%s insert {} end -text {%s} -values {%s} -open %v -tags [list %s]", el.GetID(), text, value, open, tagStr)

	nodeID := tk.Get().GetStrResult()
	node.SetID(nodeID)

	el.reference[nodeID] = node
	el.nodes = append(el.nodes, node)

	return node
}

// GetNode gets a node by its index.
func (el *TreeView) GetNode(index int) *Node {
	return el.nodes[index]
}

// GetFirstSelectedNode returns the first selected node.
// This will return nil if nothing is selected.
func (el *TreeView) GetFirstSelectedNode() *Node {
	nodes := el.GetSelectedNodes()

	if len(nodes) > 0 {
		return nodes[0]
	}

	return nil
}

// GetSelectedNodes gets all the selected nodes as an array.
func (el *TreeView) GetSelectedNodes() []*Node {
	tk.Get().Eval("%s selection", el.GetID())
	str := tk.Get().GetStrResult()

	result := make([]*Node, 0)

	if str != "" {
		for _, id := range strings.Split(str, " ") {
			if node, ok := el.reference[id]; ok {
				result = append(result, node)
			}
		}
	}

	return result
}

// Node represents a node in the tree view.
type Node struct {
	element.Ele

	nodes []*Node
}

// GetText gets the node text.
func (el *Node) GetText() string {
	tk.Get().Eval("%s item %s -text", el.GetParent().GetID(), el.GetID())

	return tk.Get().GetStrResult()
}

// GetValue gets the node value.
func (el *Node) GetValue() string {
	tk.Get().Eval("%s item %s -value", el.GetParent().GetID(), el.GetID())

	return tk.Get().GetStrResult()
}

// GetValue gets the node tags.
func (el *Node) GetTags() []string {
	tk.Get().Eval("%s item %s -tags", el.GetParent().GetID(), el.GetID())
	tagStr := tk.Get().GetStrResult()

	return strings.Split(tagStr, " ")
}

// GetNode gets a child node by its index.
func (el *Node) GetNode(index int) *Node {
	return el.nodes[index]
}

// AddNode adds a node to this node.
func (el *Node) AddNode(text, value string, open bool, tags ...string) *Node {
	node := &Node{
		nodes: make([]*Node, 0),
	}

	node.SetParent(el.GetParent())

	tagStr := strings.Join(tags, " ")
	tk.Get().Eval("%s insert %s end -text {%s} -values {%s} -open %v -tags [list %s]", el.GetParent().GetID(), el.GetID(), text, value, open, tagStr)

	nodeID := tk.Get().GetStrResult()
	node.SetID(nodeID)

	el.GetParent().(*TreeView).reference[nodeID] = node
	node.nodes = append(node.nodes, node)

	return node
}
