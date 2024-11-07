package main

import (
	"fmt"

	"github.com/odinnordico/gormat/list"
	"github.com/odinnordico/gormat/tree"
)

const squareChildren = 4

func main() {
	fmt.Println("# TREE")
	treeTest()
	fmt.Println("# LIST")
	testList()
}

func treeTest() {
	root := tree.NewRoot("root", true)
	setupChildren(root)
	fmt.Println("## With prefix")
	fmt.Print(root.Format())
	fmt.Println("## Without prefix")
	root.SetPrintPrefix(false)
	fmt.Print(root.Format())
}

func setupChildren(root *tree.Node[string]) {
	for i := 1; i <= squareChildren; i++ {
		child := tree.NewRoot(fmt.Sprintf("child-%d", i), root.PrintPrefix())
		root.AddChildren(child)

		for j := 1; j <= squareChildren; j++ {
			grandChild := tree.NewRoot(fmt.Sprintf("grand-child-%d-%d", i, j), root.PrintPrefix())
			child.AddChildren(grandChild)

			for k := 1; k <= squareChildren; k++ {
				grandGrandChild := tree.NewRoot(fmt.Sprintf("grand-grand child-%d-%d-%d", i, j, k), root.PrintPrefix())
				grandChild.AddChildren(grandGrandChild)
			}
		}
	}
}

func testList() {
	l := list.NewList[int]('>')
	l.PushBack(list.NewItem(1))
	l.PushBack(list.NewItem(2))
	l.PushBack(list.NewItem(3))
	l.PushBack(list.NewItem(4))
	l.PushBack(list.NewItem(5))
	l.PushBack(list.NewItem(6))
	l.PushBack(list.NewItem(7))
	l.PushBack(list.NewItem(8))
	l.PushBack(list.NewItem(9))
	l.PushFront(list.NewItem(0))
	fmt.Println("## Initial with Prefix")
	fmt.Println(l.Format())
	fmt.Println("## Added 10 && without prefix")
	l.PushAt(list.NewItem(10), 5)
	l.SetPrefix(0)
	fmt.Println(l.Format())
	fmt.Println("## removed 10 && with prefix")
	_ = l.PopAt(5)
	l.SetPrefix('-')
	fmt.Println(l.Format())
}
