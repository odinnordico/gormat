package list_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/odinnordico/gormat/list"
)

type testValue struct {
	Filed1 string
	Field2 int32
	Field3 []byte
}

func (tv testValue) String() string {
	return fmt.Sprintf("%s :: %d :: %s", tv.Filed1, tv.Field2, string(tv.Field3))
}

func TestItemValue(t *testing.T) {
	type args struct {
		initialValue testValue
		updatedValue testValue
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "unique case",
			args: args{
				initialValue: testValue{
					Filed1: "first",
					Field2: 1,
					Field3: []byte("value"),
				},
				updatedValue: testValue{
					Filed1: "second",
					Field2: 2,
					Field3: []byte("val"),
				},
			},
		},
	}
	for _, tt := range tests {
		i := list.NewItem(tt.args.initialValue)
		if diff := cmp.Diff(tt.args.initialValue, i.Value()); diff != "" {
			t.Errorf("NewItem() (-want, +got) = %s", diff)
		}
		i.SetValue(tt.args.updatedValue)
		if diff := cmp.Diff(tt.args.updatedValue, i.Value()); diff != "" {
			t.Errorf("NewItem() (-want, +got) = %s", diff)
		}
	}
}

func TestListPushString(t *testing.T) {
	type args struct {
		front  string
		middle string
		back   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "push strings",
			args: args{
				front:  "front",
				middle: "middle",
				back:   "back",
			},
		},
	}
	for _, tt := range tests {
		l := list.NewList[string](0)
		l.PushFront(list.NewItem(tt.args.middle))
		l.PushFront(list.NewItem(tt.args.front))
		l.PushBack(list.NewItem(tt.args.back))
		if diff := cmp.Diff(tt.args.middle, l.At(1).Value()); diff != "" {
			t.Errorf("NewList()At() (-want, +got) = %s", diff)
		}
		if i := l.At(-1); i != nil {
			t.Errorf("At(negative) want nil but got %v", i)
		}
		if i := l.At(-l.Len()); i != nil {
			t.Errorf("At(length) want nil but got %v", i)
		}
		if err := l.PushFront(nil); err == nil {
			t.Errorf("PushFront() want error but got nil")
		}
		if err := l.PushBack(nil); err == nil {
			t.Errorf("PushBack() want error but got nil")
		}
		if err := l.PushAt(nil, 0); err == nil {
			t.Errorf("PushAt() want error but got nil")
		}

		expected := []*list.Item[string]{
			list.NewItem(tt.args.front),
			list.NewItem(tt.args.middle),
			list.NewItem(tt.args.back),
		}
		comparer := func(x, y list.Item[string]) bool {
			return x.Value() == y.Value()
		}
		if diff := cmp.Diff(expected, l.Slice(), cmp.Comparer(comparer)); diff != "" {
			t.Errorf("Slice() (-want, +got) = %s", diff)
		}
	}
}

func TestListPopIntEmpty(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "pop at empty",
		},
	}
	for _, tt := range tests {
		t.Log(tt.name)
		l := list.NewList[int](0)
		if i := l.PopFront(); i != nil {
			t.Errorf("PopFront(empty) want nil but got %v", i.Value())
		}
		if i := l.PopBack(); i != nil {
			t.Errorf("PopBack(empty) want nil but got %v", i.Value())
		}
		if i := l.PopAt(0); i != nil {
			t.Errorf("PopAt(empty) want nil but got %v", i.Value())
		}
		if !l.IsEmpty() {
			t.Errorf("IsEmpty want true but got false")
		}
	}
}
func TestListPopInt(t *testing.T) {
	type args struct {
		front  int
		middle int
		back   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "push strings",
			args: args{
				front:  1,
				middle: 5,
				back:   9,
			},
		},
	}
	for _, tt := range tests {
		l := list.NewList[int](0)
		l.PushFront(list.NewItem(tt.args.front))
		l.PushAt(list.NewItem(tt.args.middle), 1)
		l.PushBack(list.NewItem(tt.args.back))
		// again but with At
		l.PushAt(list.NewItem(tt.args.front), -1)
		l.PushAt(list.NewItem(tt.args.middle), 2)
		l.PushAt(list.NewItem(tt.args.back), l.Len()+2)

		if diff := cmp.Diff(tt.args.front, l.PopFront().Value()); diff != "" {
			t.Errorf("PopFront() (-want, +got) = %s", diff)
		}
		if diff := cmp.Diff(tt.args.middle, l.PopAt(1).Value()); diff != "" {
			t.Errorf("PopAt(middle) (-want, +got) = %s", diff)
		}
		if diff := cmp.Diff(tt.args.back, l.PopBack().Value()); diff != "" {
			t.Errorf("PopBack() (-want, +got) = %s", diff)
		}
		// PopAt out of bounds
		if diff := cmp.Diff(tt.args.front, l.PopAt(-1).Value()); diff != "" {
			t.Errorf("PopFront() (-want, +got) = %s", diff)
		}
		if diff := cmp.Diff(tt.args.back, l.PopAt(l.Len()+2).Value()); diff != "" {
			t.Errorf("PopAt(middle) (-want, +got) = %s", diff)
		}

		expected := []*list.Item[int]{
			list.NewItem(tt.args.middle),
		}
		comparer := func(x, y list.Item[int]) bool {
			return x.Value() == y.Value()
		}
		if diff := cmp.Diff(expected, l.Slice(), cmp.Comparer(comparer)); diff != "" {
			t.Errorf("Slice() (-want, +got) = %s", diff)
		}
	}
}

func TestListFormatInt(t *testing.T) {
	type args struct {
		front  int
		middle int
		back   int
		prefix rune
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "format with no rune",
			args: args{
				front:  1,
				middle: 5,
				back:   9,
			},
			expected: `1
5
9`,
		},
		{
			name: "format with rune -",
			args: args{
				front:  2,
				middle: 6,
				back:   8,
				prefix: '-',
			},
			expected: `- 2
- 6
- 8`,
		},
	}
	for _, tt := range tests {
		l := list.NewList[int](-1)
		l.SetPrefix(tt.args.prefix)
		if diff := cmp.Diff(tt.args.prefix, l.Prefix()); diff != "" {
			t.Errorf("Prefix() (-want, +got) = %s", diff)
		}
		l.PushFront(list.NewItem(tt.args.front))
		l.PushAt(list.NewItem(tt.args.middle), 1)
		l.PushBack(list.NewItem(tt.args.back))
		if diff := cmp.Diff(tt.expected, l.Format()); diff != "" {
			t.Errorf("Format() (-want, +got) = %s", diff)
		}
	}
}
