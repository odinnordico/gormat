package list

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/odinnordico/gormat/format"
)

const nilItemMsg = "the item to be added cannot be nil"

// Item is the basic element of the list
type Item[T any] struct {
	value T // value of the item
}

// SetValue sets the value of the Item
func (i *Item[T]) SetValue(v T) {
	i.value = v
}

// Value retrieves the value of the Item
func (i *Item[T]) Value() T {
	return i.value
}

// NewItem creates a new Item of the given type
func NewItem[T any](v T) *Item[T] {
	return &Item[T]{value: v}
}

// List is a structure of a sequence of Items
type List[T any] struct {
	items  []*Item[T] // items is the slices in charge of hold the Items of the List
	prefix rune       // prefix is the rune to be used when formatting the List. If prefix is 0, no prefix is printed
}

// SetPrefix updates the prefix of the list
func (l *List[T]) SetPrefix(p rune) {
	l.prefix = p
}

// Prefix retrieves the prefix of the list
func (l *List[T]) Prefix() rune {
	return l.prefix
}

// PushFront adds an item to the list in the first position. If null is passed then it returns error
func (l *List[T]) PushFront(i *Item[T]) error {
	if i == nil {
		return errors.New(nilItemMsg)
	}
	l.items = append([]*Item[T]{i}, l.items...)
	return nil
}

// PushBack adds an item to the end of the list. If null is passed then it returns error
func (l *List[T]) PushBack(i *Item[T]) error {
	if i == nil {
		return errors.New(nilItemMsg)
	}
	l.items = append(l.items, i)
	return nil
}

// PushAt adds an item at a given index. If null is passed then it returns error
// If the index is greater than the length of the list then it is pushed at the back
// If the index is less than zero then it is pushed at the front
func (l *List[T]) PushAt(i *Item[T], idx int) error {
	if i == nil {
		return errors.New(nilItemMsg)
	}
	if idx >= len(l.items) {
		return l.PushBack(i)
	}
	if idx <= 0 {
		return l.PushFront(i)
	}

	l.items = append(l.items[:idx], append([]*Item[T]{i}, l.items[idx:]...)...)
	return nil
}

// PopFront deletes the first Item of the List and it is returned
func (l *List[T]) PopFront() *Item[T] {
	if len(l.items) == 0 {
		return nil
	}
	i := l.items[0]
	l.items = l.items[1:] // Remove the first item by slicing.
	return i
}

// PopBack deletes the last Item of the List and it is returned
func (l *List[T]) PopBack() *Item[T] {
	if len(l.items) == 0 {
		return nil
	}
	idx := len(l.items) - 1
	i := l.items[idx]
	l.items = l.items[:idx]
	return i
}

// PopAt deletes the Item at the given index and it is returned
// If the index is greater than the length of the List the last Item is returned
// If the index is less than zero then the first Item is returned
func (l *List[T]) PopAt(idx int) *Item[T] {
	if len(l.items) == 0 {
		return nil
	}
	if idx >= len(l.items) {
		return l.PopBack()
	}
	if idx <= 0 {
		return l.PopFront()
	}
	i := l.items[idx]
	l.items = append(l.items[:idx], l.items[idx+1:]...)
	return i
}

// At return the Item at the given index keeping it in the List
// If the index is greater than the length of the List, nil is returned
// If the index is less than zero, nil is returned
func (l *List[T]) At(i int) *Item[T] {
	if i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Len return the number of Items in the List
func (l *List[T]) Len() int {
	return len(l.items)
}

// Slice return the list as an slice.
// Be aware that the items returned are the pointers kept in the List
func (l *List[T]) Slice() []*Item[T] {
	return l.items
}

// IsEmpty returns true if the list has zero Items, otherwise returns false
func (l *List[T]) IsEmpty() bool {
	return len(l.items) == 0
}

// Format transforms the Items in List to a ordered string separated each by new line
// if the prefix is set to be greater than zero, the Items will be written with that rune followed by a space
func (l *List[T]) Format() string {
	var builder strings.Builder
	if l.prefix > 0 {
		builder.WriteRune(l.prefix)
		builder.WriteString(" ")
	}

	for idx, item := range l.items {
		if idx > 0 {
			builder.WriteString(format.NewLine)
			if l.prefix != 0 {
				builder.WriteRune(l.prefix)
				builder.WriteString(" ")
			}
		}
		builder.WriteString(fmt.Sprintf("%v", item.value))
	}
	return builder.String()
}

// NewList creates a List with a given prefix
func NewList[T any](prefix rune) *List[T] {
	if !unicode.IsPrint(prefix) {
		prefix = 0
	}
	return &List[T]{prefix: prefix}
}
