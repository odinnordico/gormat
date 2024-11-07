package list

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/odinnordico/gormat/format"
)

const nilItemMsg = "the item to be added cannot be nil"

type Item[T any] struct {
	value T
}

func (i *Item[T]) SetValue(v T) {
	i.value = v
}

func (i *Item[T]) Value() T {
	return i.value
}

func NewItem[T any](v T) *Item[T] {
	return &Item[T]{value: v}
}

type List[T any] struct {
	items  []*Item[T]
	prefix rune
}

func (l *List[T]) SetPrefix(p rune) {
	l.prefix = p
}

func (l *List[T]) Prefix() rune {
	return l.prefix
}

func (l *List[T]) PushFront(i *Item[T]) error {
	if i == nil {
		return errors.New(nilItemMsg)
	}
	l.items = append([]*Item[T]{i}, l.items...)
	return nil
}

func (l *List[T]) PushBack(i *Item[T]) error {
	if i == nil {
		return errors.New(nilItemMsg)
	}
	l.items = append(l.items, i)
	return nil
}

func (l *List[T]) PushAt(i *Item[T], idx int) error {
	if i == nil {
		return errors.New(nilItemMsg)
	}
	if idx >= len(l.items) {
		return l.PushBack(i)
	}
	l.items = append(l.items[:idx], append([]*Item[T]{i}, l.items[idx:]...)...)
	return nil
}

func (l *List[T]) PopFront() *Item[T] {
	if len(l.items) == 0 {
		return nil
	}
	i := l.items[0]
	l.items = l.items[1:] // Remove the first item by slicing.
	return i
}

func (l *List[T]) PopBack() *Item[T] {
	if len(l.items) == 0 {
		return nil
	}
	idx := len(l.items) - 1
	i := l.items[idx]
	l.items = l.items[:idx]
	return i
}

func (l *List[T]) PopAt(idx int) *Item[T] {
	if len(l.items) == 0 || idx >= len(l.items) {
		return nil
	}
	i := l.items[idx]
	l.items = append(l.items[:idx], l.items[idx+1:]...)
	return i
}

func (l *List[T]) At(i int) *Item[T] {
	if i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

func (l *List[T]) Len() int {
	return len(l.items)
}

func (l *List[T]) Slice() []*Item[T] {
	return l.items
}

func (l *List[T]) IsEmpty() bool {
	return len(l.items) == 0
}

func (l *List[T]) Format() string {
	var builder strings.Builder
	if l.prefix != 0 {
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

func NewList[T any](prefix rune) *List[T] {
	if !unicode.IsPrint(prefix) {
		prefix = 0
	}
	return &List[T]{prefix: prefix}
}
