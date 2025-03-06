package tree

import (
	"bytes"
	"fmt"

	"github.com/odinnordico/gormat/format"
)

// Prefix defines the common tree prefixes
type Prefix string

const (
	firstElemPrefix Prefix = `├─ ` // Prefix for the first branch element
	lastElemPrefix  Prefix = `└─ ` // Prefix for the last branch element
	space           Prefix = "   " // Prefix blank space
	pipe            Prefix = `│  ` // Prefix for non element
)

// Node is the basic element of a tree
type Node[T any] struct {
	value       T          // value of the node
	parent      *Node[T]   // parent of the current node
	children    []*Node[T] // children of the current node
	printPrefix bool       // pintPrefix indicates if the Prefix will be used  when formatting the Tree. It is only necessary to be set in the root

}

// SetValue sets the value of the Node
func (n *Node[T]) SetValue(v T) {
	n.value = v
}

// Value retrieves the value of the Node
func (n *Node[T]) Value() T {
	return n.value
}

// PrintPrefix checks whether if print the prefix or not
func (n *Node[T]) PrintPrefix() bool {
	return n.printPrefix
}

// SetPrintPrefix sets the flag to print the prefix of the Node and its children
func (n *Node[T]) SetPrintPrefix(p bool) {
	n.printPrefix = p
	if !n.IsLeaf() {
		for i := range n.children {
			n.children[i].SetPrintPrefix(p)
		}
	}
}

// Format generates a string representation of the tree with visual hierarchy.
func (n *Node[T]) Format() string {
	var buffer bytes.Buffer
	n.formatHelper(&buffer, "", true)
	return buffer.String()
}

// formatHelper is a recursive helper to generate the tree structure.
func (n *Node[T]) formatHelper(buffer *bytes.Buffer, prefix string, isLast bool) {
	val := format.CleanString(fmt.Sprintf("%v", n.value))
	// Choose prefix based on whether this node is the last child.
	if n.IsRoot() {
		buffer.WriteString(format.CleanString(val))
	} else {
		branchPrefix := lastElemPrefix
		if !isLast {
			branchPrefix = firstElemPrefix
		}
		if !n.printPrefix {
			branchPrefix = space
		}
		fmt.Fprintf(buffer, "%s%s%s", prefix, branchPrefix, val)
	}
	buffer.WriteString(format.NewLine)

	// Update the prefix for child nodes
	childPrefix := prefix
	if !n.IsRoot() {
		if isLast || !n.printPrefix {
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
	c.printPrefix = n.printPrefix
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

// NewNode creates a root Node with a given value and if prefix print
// The Node is root because it does not has either father nor children
// but when added as a child the father is being updated
func NewNode[T any](v T, prefix bool) *Node[T] {
	return &Node[T]{
		value:       v,
		printPrefix: prefix,
	}
}
