package tree

import (
	"bytes"
	"fmt"

	"github.com/odinnordico/gormat/format"
)

type Prefix string

const (
	firstElemPrefix Prefix = `├─ `
	lastElemPrefix  Prefix = `└─ `
	space           Prefix = "   "
	pipe            Prefix = `│  `
	newLine                = "\n\r"
)

type Node[T any] struct {
	value    T
	parent   *Node[T]
	children []*Node[T]
}

func (n *Node[T]) SetValue(v T) {
	n.value = v
}

func (n *Node[T]) Value() T {
	return n.value
}

// Format generates a string representation of the tree with visual hierarchy.
func (n *Node[T]) Format() string {
	var buffer bytes.Buffer
	n.formatHelper(&buffer, "", true)
	return buffer.String()
}

// formatHelper is a recursive helper to generate the tree structure.
func (n *Node[T]) formatHelper(buffer *bytes.Buffer, prefix string, isLast bool) {
	// Choose prefix based on whether this node is the last child.
	if n.IsRoot() {
		buffer.WriteString(format.CleanString(fmt.Sprintf("%v", n.value)))
	} else {
		branchPrefix := lastElemPrefix
		if !isLast {
			branchPrefix = firstElemPrefix
		}
		buffer.WriteString(fmt.Sprintf("%s%s%s", prefix, branchPrefix, format.CleanString(fmt.Sprintf("%v", n.value))))
	}
	buffer.WriteString("\n")

	// Update the prefix for child nodes
	childPrefix := prefix
	if !n.IsRoot() {
		if isLast {
			childPrefix += string(space)
		} else {
			childPrefix += string(pipe)
		}
	}

	// Recursively format each child node.
	for i, child := range n.children {
		child.formatHelper(buffer, childPrefix, i == len(n.children)-1)
	}
}

// IsRoot checks if a node is the root.
func (n *Node[T]) IsRoot() bool {
	return n.parent == nil
}

// IsLeaf checks if a node is a leaf node.
func (n *Node[T]) IsLeaf() bool {
	return len(n.children) == 0
}

// addChild safely adds a child node and sets its parent to the current node.
func (n *Node[T]) addChild(c *Node[T]) {
	c.parent = n
	n.children = append(n.children, c)
}

// AddChildren adds multiple children at once.
func (n *Node[T]) AddChildren(c ...*Node[T]) {
	for i := range c {
		n.addChild(c[i])
	}
}

// Level calculates the depth level of a node.
func (n *Node[T]) Level() int {
	if n.IsRoot() {
		return 1
	}
	return 1 + n.parent.Level()
}

func (n *Node[T]) Children() []*Node[T] {
	return n.children
}

// NewRoot creates a new root node with a given value.
func NewRoot[T any](v T) *Node[T] {
	return &Node[T]{value: v}
}
