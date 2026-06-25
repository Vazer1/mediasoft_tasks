package main

import "fmt"

type doublyLinkedList struct {
	first *item
	last  *item
	size  int
}

type item struct {
	v    any
	next *item
	prev *item
}

func newDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{}
}

// add - добавление значения в связный список
func add(l *doublyLinkedList, v any) {
	curr := item{v: v}
	if l.last == nil {
		l.first = &curr
		l.last = &curr
	} else {
		curr.prev = l.last
		l.last.next = &curr
		l.last = &curr
	}
	l.size++
}

// get - получение значения по индексу из связанного списка
func get(l *doublyLinkedList, idx int) any {
	if l.first == nil || idx < 0 || idx >= l.size {
		return nil
	}
	curr := l.first
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	return curr.v
}

// remove - удаление значения по индексу из списка
func remove(l *doublyLinkedList, idx int) {
	if l.first == nil || idx < 0 || idx >= l.size {
		return
	}

	curr := l.first
	for i := 0; i < idx; i++ {
		curr = curr.next
	}

	if curr.prev == nil {
		l.first = curr.next
	} else {
		curr.prev.next = curr.next
	}

	if curr.next == nil {
		l.last = curr.prev
	} else {
		curr.next.prev = curr.prev
	}
	l.size--
}

// values - получение слайса значений из списка
func values(l *doublyLinkedList) []any {
	res := make([]any, l.size)
	if l.first == nil {
		return res
	}
	curr := l.first
	for i := 0; i < l.size; i++ {
		res[i] = curr.v
		curr = curr.next
	}
	return res
}

var commands = `
1 - add
2 - get
3 - remove
4 - values
5 - Выйти из программы
`

func main() {
	myList := newDoublyLinkedList()

	for {
		var val int
		cmd := 0
		fmt.Print(commands)
		fmt.Scanln(&cmd)

		switch cmd {
		case 1:
			fmt.Println("Введите число")
			fmt.Scanln(&val)
			add(myList, val)
		case 2:
			fmt.Println("Введите индекс")
			fmt.Scanln(&val)
			val := get(myList, val)
			fmt.Println(val)
		case 3:
			fmt.Println("Введите индекс")
			fmt.Scanln(&val)
			remove(myList, val)
		case 4:
			slice := values(myList)
			fmt.Println(slice)
		case 5:
			return
		}
	}
}