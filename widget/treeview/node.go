package treeview

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

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

// GetOpen gets whether the node is open.
func (el *Node) GetOpen() bool {
	tk.Get().Eval("%s item %s -open", el.GetParent().GetID(), el.GetID())

	return tk.Get().GetBoolResult()
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

	el.GetParent().(*TreeView).nodeRef[nodeID] = node
	node.nodes = append(node.nodes, node)

	return node
}
