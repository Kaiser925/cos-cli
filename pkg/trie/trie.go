package trie

type Tree[T any] struct {
	Children map[string]*Tree[T]
	Value    T

	segment StringSegmentFunc
}

func New[T any](seg StringSegmentFunc) *Tree[T] {
	var none T
	return &Tree[T]{
		Value:    none,
		Children: make(map[string]*Tree[T]),
		segment:  seg,
	}
}

func (t *Tree[T]) Get(name string) (*Tree[T], bool) {
	node := t
	for part, i := t.segment(name, 0); part != ""; part, i = t.segment(name, i) {
		node = node.Children[part]
		if node == nil {
			return nil, false
		}
	}
	return node, true
}

func (t *Tree[T]) Put(name string, val T) {
	node := t
	for part, i := t.segment(name, 0); part != ""; part, i = t.segment(name, i) {
		child, ok := node.Children[part]
		if !ok {
			child = New[T](t.segment)
			node.Children[part] = child
		}
		node = child
	}
	node.Value = val
}
