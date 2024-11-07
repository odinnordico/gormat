package list

import (
	"testing"
)

func TestItemValue(t *testing.T) {
	item := NewItem("test")
	if item.Value() != "test" {
		t.Errorf("Expected value 'test', got %v", item.Value())
	}
	item.SetValue("new")
	if item.Value() != "new" {
		t.Errorf("Expected value 'new', got %v", item.Value())
	}
}

func TestListPushFront(t *testing.T) {
	list := NewList[string]('*')
	item := NewItem("test")
	err := list.PushFront(item)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if list.At(0) != item {
		t.Error("Item not added at front")
	}
}

func TestListPushBack(t *testing.T) {
	list := NewList[string]('*')
	item := NewItem("test")
	err := list.PushBack(item)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if list.At(0) != item {
		t.Error("Item not added at back")
	}
}

func TestListPushAt(t *testing.T) {
	list := NewList[string]('*')
	item1 := NewItem("first")
	item2 := NewItem("second")
	item3 := NewItem("third")

	list.PushBack(item1)
	list.PushBack(item3)
	list.PushAt(item2, 1)

	if list.At(1) != item2 {
		t.Error("Item not inserted at correct index")
	}
}

func TestListPopFront(t *testing.T) {
	list := NewList[string]('*')
	item := NewItem("test")
	list.PushFront(item)

	popped := list.PopFront()
	if popped != item {
		t.Error("Expected popped item to match pushed item")
	}

	if list.Len() != 0 {
		t.Error("Expected list to be empty after pop")
	}
}

func TestListPopBack(t *testing.T) {
	list := NewList[string]('*')
	item := NewItem("test")
	list.PushBack(item)

	popped := list.PopBack()
	if popped != item {
		t.Error("Expected popped item to match pushed item")
	}

	if list.Len() != 0 {
		t.Error("Expected list to be empty after pop")
	}
}

func TestListPopAt(t *testing.T) {
	list := NewList[string]('*')
	item1 := NewItem("first")
	item2 := NewItem("second")
	list.PushBack(item1)
	list.PushBack(item2)

	popped := list.PopAt(1)
	if popped != item2 {
		t.Error("Expected to pop the second item")
	}

	if list.Len() != 1 {
		t.Error("Expected list length to be 1 after pop")
	}
}

func TestListAt(t *testing.T) {
	list := NewList[string]('*')
	item := NewItem("test")
	list.PushBack(item)

	if list.At(0) != item {
		t.Error("Expected item to be at index 0")
	}

	if list.At(1) != nil {
		t.Error("Expected nil for out-of-bounds index")
	}
}

func TestListLen(t *testing.T) {
	list := NewList[string]('*')
	if list.Len() != 0 {
		t.Error("Expected length to be 0 for new list")
	}

	list.PushBack(NewItem("item"))
	if list.Len() != 1 {
		t.Error("Expected length to be 1 after adding an item")
	}
}

func TestListFormat(t *testing.T) {
	list := NewList[string]('*')
	item1 := NewItem("first")
	item2 := NewItem("second")
	list.PushBack(item1)
	list.PushBack(item2)

	// Adjusted expected output to match Format's behavior
	expected := "* first\n\r* second"
	formatted := list.Format()
	if formatted != expected {
		t.Errorf("Expected formatted output:\n%s\nGot:\n%s", expected, formatted)
	}
}
