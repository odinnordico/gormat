package main

import (
	"fmt"

	"github.com/odinnordico/gormat/tree"
)

const squareChildren = 4

func main() {
	root := tree.NewRoot("root")
	setupChildren(root)
	fmt.Print(root.Format())
}

func setupChildren(root *tree.Node[string]) {
	for i := 1; i <= squareChildren; i++ {
		child := tree.NewRoot(fmt.Sprintf("child-%d", i))
		root.AddChildren(child)

		for j := 1; j <= squareChildren; j++ {
			grandChild := tree.NewRoot(fmt.Sprintf("grand-child-%d-%d", i, j))
			child.AddChildren(grandChild)

			for k := 1; k <= squareChildren; k++ {
				grandGrandChild := tree.NewRoot(fmt.Sprintf("grand-grand child-%d-%d-%d", i, j, k))
				grandChild.AddChildren(grandGrandChild)
			}
		}
	}
}
