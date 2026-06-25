package main

import "fmt"

type tree struct {
	head *node
}

type node struct {
	left, right *node
	v           int
}

func newTree() *tree {
	return &tree{}
}

// add - добавление значения в дерево
func add(t *tree, v int) {
	newNode := &node{v: v}
	if t.head == nil {
		t.head = newNode
		return
	}
	curr := t.head
	if curr == nil {
		return
	}
	for {
		if v < curr.v {
			if curr.left == nil {
				curr.left = newNode
				return
			} else {
				curr = curr.left
			}
		} else if v > curr.v {
			if curr.right == nil {
				curr.right = newNode
				return
			} else {
				curr = curr.right
			}
		} else {
			return
		}
	}
}

// remove - удаление значения из дерева
func remove(l *tree, v int) {
	var parent *node = nil
	curr := l.head
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

	var child *node
	if curr.left != nil {
		child = curr.left
	} else {
		child = curr.right
	}

	if parent == nil {
		l.head = child
	} else if parent.left == curr {
		parent.left = child
	} else {
		parent.right = child
	}
}

// values - получение отсортированного слайса значений из дерева
func values(l *tree) []int {
	var result []int

	var recursion func(n *node)
	recursion = func(n *node) {
		if n == nil {
			return
		}
		recursion(n.left)
		result = append(result, n.v)
		recursion(n.right)
	}

	recursion(l.head)
	return result
}

var commands = `
1 - add
2 - remove
3 - values
4 - Выйти из программы
`

func main() {
	myTree := newTree()

	for {
		var val int
		cmd := 0
		fmt.Print(commands)
		fmt.Scanln(&cmd)

		switch cmd {
		case 1:
			fmt.Println("Введите число")
			fmt.Scanln(&val)
			add(myTree, val)
		case 2:
			fmt.Println("Введите число")
			fmt.Scanln(&val)
			remove(myTree, val)
		case 3:
			slice := values(myTree)
			fmt.Println(slice)
		case 4:
			return
		}
	}
}