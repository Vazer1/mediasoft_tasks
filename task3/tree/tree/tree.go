package tree

import "cmp"

type node[T cmp.Ordered] struct {
	left, right *node[T]
	v           T
}

type Tree[T cmp.Ordered] struct {
	head *node[T]
}

func New[T cmp.Ordered]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) Add(v T) {
	newNode := &node[T]{v: v}
	if t.head == nil {
		t.head = newNode
		return
	}
	
	curr := t.head
	for {
		if v < curr.v {
			if curr.left == nil {
				curr.left = newNode
				return
			}
			curr = curr.left
		} else if v > curr.v {
			if curr.right == nil {
				curr.right = newNode
				return
			}
			curr = curr.right
		} else {
			return
		}
	}
}

func (t *Tree[T]) Remove(v T) {
	var parent *node[T] = nil
	curr := t.head
	
	for curr != nil && curr.v != v {
		parent = curr
		if v < curr.v {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	if curr == nil {
		return
	} else if curr.left != nil && curr.right != nil {
		replacementParent := curr
		replacement := curr.right
		for replacement.left != nil {
			replacementParent = replacement
			replacement = replacement.left
		}
		curr.v = replacement.v
		curr = replacement
		parent = replacementParent
	}

	var child *node[T]
	if curr.left != nil {
		child = curr.left
	} else {
		child = curr.right
	}

	if parent == nil {
		t.head = child
	} else if parent.left == curr {
		parent.left = child
	} else {
		parent.right = child
	}
}

func (t *Tree[T]) Values() []T {
	var result []T

	var recursion func(n *node[T])
	recursion = func(n *node[T]) {
		if n == nil {
			return
		}
		recursion(n.left)
		result = append(result, n.v)
		recursion(n.right)
	}

	recursion(t.head)
	return result
}