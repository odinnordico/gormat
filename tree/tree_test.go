package tree_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/odinnordico/gormat/tree"
)

type testValue struct {
	Filed1 string
	Field2 int32
	Field3 []byte
}

func (tv testValue) String() string {
	return fmt.Sprintf("%s :: %d :: %s", tv.Filed1, tv.Field2, string(tv.Field3))
}

func TestNodeValue(t *testing.T) {
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
		r := tree.NewNode(tt.args.initialValue, false)
		if diff := cmp.Diff(tt.args.initialValue, r.Value()); diff != "" {
			t.Errorf("NewNode() (-want, +got) = %s", diff)
		}
		r.SetValue(tt.args.updatedValue)
		if diff := cmp.Diff(tt.args.updatedValue, r.Value()); diff != "" {
			t.Errorf("NewNode() (-want, +got) = %s", diff)
		}
		if diff := cmp.Diff(true, r.IsLeaf()); diff != "" {
			t.Errorf("IsLeaf() (-want, +got) = %s", diff)
		}
		if diff := cmp.Diff(true, r.IsRoot()); diff != "" {
			t.Errorf("IsRoot() (-want, +got) = %s", diff)
		}
		// check child prefix
		c := tree.NewNode(tt.args.initialValue, false)
		r.AddChildren(c)
		if diff := cmp.Diff(false, r.IsLeaf()); diff != "" {
			t.Errorf("IsLeaf(updated) (-want, +got) = %s", diff)
		}
		if diff := cmp.Diff(true, c.IsLeaf()); diff != "" {
			t.Errorf("IsLeaf(child) (-want, +got) = %s", diff)
		}
		r.SetPrintPrefix(true)
		if diff := cmp.Diff(true, r.PrintPrefix()); diff != "" {
			t.Errorf("PrintPrefix() (-want, +got) = %s", diff)
		}
		// Level
		if diff := cmp.Diff(1, r.Level()); diff != "" {
			t.Errorf("Level(parent) (-want, +got) = %s", diff)
		}
		if diff := cmp.Diff(2, c.Level()); diff != "" {
			t.Errorf("Level(child) (-want, +got) = %s", diff)
		}
		// Children
		comparator := func(x, y *tree.Node[testValue]) bool {
			return x.Value().String() == y.Value().String()
		}
		if diff := cmp.Diff([]*tree.Node[testValue]{c}, r.Children(), cmp.Comparer(comparator)); diff != "" {
			t.Errorf("Children() (-want, +got) = %s", diff)
		}

	}
}

func TestTreeFormat(t *testing.T) {
	//TODO
}
