package tree

import (
	"testing"
)

func TestNodeSetValue(t *testing.T) {
	n := &Node[string]{}
	n.SetValue("root")
	if n.Value() != "root" {
		t.Errorf("Expected value to be 'root', got %v", n.Value())
	}
}

func TestNodeAddChildren(t *testing.T) {
	n := &Node[string]{}
	c1 := &Node[string]{}
	c2 := &Node[string]{}

	n.AddChildren(c1, c2)

	if len(n.Children()) != 2 {
		t.Errorf("Expected 2 children, got %d", len(n.Children()))
	}

	if n.Children()[0] != c1 || n.Children()[1] != c2 {
		t.Error("Children not added correctly")
	}
}

func TestNodeIsRoot(t *testing.T) {
	n := &Node[string]{}
	c := &Node[string]{}
	n.AddChildren(c)

	if !n.IsRoot() {
		t.Error("Expected root node")
	}

	if c.IsRoot() {
		t.Error("Expected child node not to be root")
	}
}

func TestNodeIsLeaf(t *testing.T) {
	n := &Node[string]{}
	c := &Node[string]{}
	n.AddChildren(c)

	if n.IsLeaf() {
		t.Error("Expected parent node not to be leaf")
	}

	if !c.IsLeaf() {
		t.Error("Expected child node to be leaf")
	}
}

func TestNodeLevel(t *testing.T) {
	root := &Node[string]{}
	child := &Node[string]{}
	grandChild := &Node[string]{}

	root.AddChildren(child)
	child.AddChildren(grandChild)

	if root.Level() != 1 {
		t.Errorf("Expected root level to be 1, got %d", root.Level())
	}

	if child.Level() != 2 {
		t.Errorf("Expected child level to be 2, got %d", child.Level())
	}

	if grandChild.Level() != 3 {
		t.Errorf("Expected grandchild level to be 3, got %d", grandChild.Level())
	}
}

func TestNodeFormat(t *testing.T) {
	root := &Node[string]{}
	root.SetValue("root")
	child1 := &Node[string]{}
	child1.SetValue("child1")
	child2 := &Node[string]{}
	child2.SetValue("child2")

	root.AddChildren(child1, child2)
	grandChild := &Node[string]{}
	grandChild.SetValue("grandChild")
	child1.AddChildren(grandChild)

	formatted := root.Format()
	expected := "root\n\r   child1\n\r      grandChild\n\r   child2\n\r"

	if formatted != expected {
		t.Errorf("Expected formatted output:\n%s\nGot:\n%s", expected, formatted)
	}
}

func TestNodeAddChild(t *testing.T) {
	parent := &Node[string]{}
	child := &Node[string]{}

	parent.addChild(child)

	if len(parent.Children()) != 1 {
		t.Errorf("Expected 1 child, got %d", len(parent.Children()))
	}

	if parent.Children()[0] != child {
		t.Error("Child not added correctly")
	}
}
