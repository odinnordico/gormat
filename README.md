# Gormat

[![Go Report Card](https://goreportcard.com/badge/github.com/odinnordico/gormat)](https://goreportcard.com/report/github.com/odinnordico/gormat)

**Gormat** is a Go library that provides custom implementations of a list and tree data structure with support for formatting and string manipulation utilities. These structures offer a flexible way to manage and display hierarchical or linear collections of data in Go.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Installation](#installation)
3. [Usage](#usage)
   - [List](#list)
   - [Tree](#tree)
4. [API Documentation](#api-documentation)
5. [Formatting Options](#formatting-options)
6. [Testing](#testing)
7. [Examples](#examples)
8. [Contributing](#contributing)
9. [License](#license)

---

## Getting Started

### Prerequisites

Make sure you have Go installed (version 1.18 or higher is recommended). You can download Go from [here](https://golang.org/dl/).

### Installation

To add this library to your project, you can use the following command:

```bash
go get github.com/odinnordico/gormat
```

This will download the `gormat` package and make it available in your project.

## Usage

### Importing the Library

The library is split into three main packages:

1. `list` for handling list operations.
2. `tree` for handling tree-based hierarchical structures.
3. `format` for utility functions to clean and format strings.

To use any of these, you need to import them into your Go files:

```go
import (
    "github.com/odinnordico/gormat/list"
    "github.com/odinnordico/gormat/tree"
    "github.com/odinnordico/gormat/format"
)
```

### List

The `list` package provides a linked list-like data structure where you can add, remove, and manipulate elements with custom prefixes.

#### Creating a List

```go
l := list.NewList[string]('*') // Creates a new list with '*' prefix
```

#### Adding Items to the List

```go
item1 := list.NewItem("first item")
item2 := list.NewItem("second item")

// Push to front
l.PushFront(item1)

// Push to back
l.PushBack(item2)

// Insert at a specific index
l.PushAt(item2, 1)
```

#### Removing Items from the List

```go
// Remove the first item
removedItem := l.PopFront()

// Remove the last item
removedItem = l.PopBack()

// Remove an item at a specific index
removedItem = l.PopAt(1)
```

#### Retrieving and Formatting Items

```go
// Retrieve an item by index
item := l.At(0)

// Get the list length
length := l.Len()

// Format the list with the prefix
formattedOutput := l.Format()
fmt.Println(formattedOutput)
```

### Tree

The `tree` package provides a simple, generic tree structure with support for hierarchical data representation.

#### Creating a Tree

```go
root := &tree.Node[string]{}
root.SetValue("Root")
```

#### Adding Children

```go
child1 := &tree.Node[string]{}
child1.SetValue("Child 1")
root.AddChildren(child1)
```

#### Formatting the Tree

The `Format` function in the `tree` package provides a visual representation of the tree hierarchy.

```go
fmt.Println(root.Format())
```

### String Formatting Utilities

The `format` package provides a `CleanString` function to remove non-printable and non-graphic characters from strings.

```go
cleaned := format.CleanString("some string with ​ non-printable characters")
```

## API Documentation

### list Package

#### `type Item[T any]`

Represents an individual item within the list.

- `func NewItem[T any](v T) *Item[T]`: Creates a new list item.
- `func (i *Item[T]) SetValue(v T)`: Sets the value of an item.
- `func (i *Item[T]) Value() T`: Gets the value of an item.

#### `type List[T any]`

Represents the list data structure.

- `func NewList[T any](prefix rune) *List[T]`: Creates a new list with an optional prefix.
- `func (l *List[T]) SetPrefix(p rune)`: Sets the list prefix.
- `func (l *List[T]) Prefix() rune`: Gets the list prefix.
- `func (l *List[T]) PushFront(i *Item[T]) error`: Adds an item to the front of the list.
- `func (l *List[T]) PushBack(i *Item[T]) error`: Adds an item to the back of the list.
- `func (l *List[T]) PushAt(i *Item[T], idx int) error`: Adds an item at a specified index.
- `func (l *List[T]) PopFront() *Item[T]`: Removes an item from the front of the list.
- `func (l *List[T]) PopBack() *Item[T]`: Removes an item from the back of the list.
- `func (l *List[T]) PopAt(idx int) *Item[T]`: Removes an item at a specific index.
- `func (l *List[T]) At(i int) *Item[T]`: Retrieves an item by index.
- `func (l *List[T]) Len() int`: Returns the list length.
- `func (l *List[T]) Slice() []*Item[T]`: Returns the list as a slice.
- `func (l *List[T]) Format() string`: Formats the list as a string.

### tree Package

#### `type Node[T any]`

Represents a single node in the tree.

- `func (n *Node[T]) SetValue(v T)`: Sets the value of the node.
- `func (n *Node[T]) Value() T`: Gets the value of the node.
- `func (n *Node[T]) Format() string`: Formats the node and its children.
- `func (n *Node[T]) IsRoot() bool`: Checks if the node is the root.
- `func (n *Node[T]) IsLeaf() bool`: Checks if the node is a leaf.
- `func (n *Node[T]) AddChildren(c ...*Node[T])`: Adds children to the node.
- `func (n *Node[T]) Children() []*Node[T]`: Retrieves all children of the node.

### format Package

#### `func CleanString(s string) string`

Cleans a string by removing non-graphic and non-printable characters and trimming whitespace.

## Testing

### Running Tests

The library includes extensive unit tests for both the `list` and `tree` packages. You can run all tests using the following command:

```bash
go test ./... -cover
```

This command will run all tests in the library and show coverage.

### Adding Test Coverage

If you'd like to add more tests, make sure to place them in `*_test.go` files within the relevant packages and use Go’s testing package.

### Example Test Files

- `list_test.go`: Tests for the `List` and `Item` functionalities.
- `tree_test.go`: Tests for the `Tree` and `Node` functionalities.

## Examples

### Basic List Example

```go
package main

import (
    "fmt"
    "github.com/odinnordico/gormat/list"
)

func main() {
    myList := list.NewList[string]('*')
    myList.PushBack(list.NewItem("First"))
    myList.PushBack(list.NewItem("Second"))

    fmt.Println(myList.Format()) // Output should include * prefix
}
```

### Basic Tree Example

```go
package main

import (
    "fmt"
    "github.com/odinnordico/gormat/tree"
)

func main() {
    root := &tree.Node[string]{}
    root.SetValue("Root")

    child := &tree.Node[string]{}
    child.SetValue("Child")
    root.AddChildren(child)

    fmt.Println(root.Format())
}
```

## Contributing

We welcome contributions! If you'd like to improve Gormat, please fork the repository, create a new branch, and submit a pull request. Be sure to run tests before submitting.

### Guidelines

- Follow idiomatic Go style.
- Write tests for any new functionality.
- Document your code.

## License

Gormat is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

This README provides a thorough guide to understanding, installing, and using the `gormat` library. If you have questions, feel free to open an issue on the GitHub repository.
