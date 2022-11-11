package trie

import "testing"

var nodes = []struct {
	path string
	val  string
}{
	{"a", "a"},
	{"/child", "child"},
	{"/child/child2", "child2"},
	{"/child/child2/child3", "child3"},
}

func TestTree_Get(t *testing.T) {
	root := New[string](PathSegment)
	for _, node := range nodes {
		root.Put(node.path, node.val)
	}
	nones := []string{"/path/not", "/", "/dasd/cdac"}
	for _, test := range nones {
		if _, ok := root.Get(test); ok {
			t.Errorf("Get(%v) failed, should be false", test)
		}
	}

	for _, node := range nodes {
		if val, _ := root.Get(node.path); val.Value != node.val {
			t.Errorf("Get(%v) failed, want %v got %v", node.val, node.val, val)
		}
	}
}
