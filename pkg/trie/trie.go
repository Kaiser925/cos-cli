package trie

type Tree[T any] struct {
	children map[string]*Tree[T]
	segment  StringSegmentFunc

	value T
}

func New[T any](val T, seg StringSegmentFunc) *Tree[T] {
	return &Tree[T]{
		value:    val,
		children: make(map[string]*Tree[T]),
		segment:  seg,
	}
}

func (t *Tree[T]) Get(name string) (T, bool) {
	node := t
	for part, i := t.segment(name, 0); part != ""; part, i = t.segment(name, i) {
		node = node.children[part]
		if node == nil {
			var none T
			return none, false
		}
	}
	return node.value, true
}

func (t *Tree[T]) Put(name string, val T) {
	node := t
	for part, i := t.segment(name, 0); part != ""; part, i = t.segment(name, i) {
		child, ok := node.children[part]
		if !ok {
			var none T
			child = New[T](none, t.segment)
			node.children[part] = child
		}
		node = child
	}
	node.value = val
}
