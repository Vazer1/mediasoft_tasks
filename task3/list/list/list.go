package list

type item[T any] struct {
	v    T
	next *item[T]
	prev *item[T]
}

type DoublyLinkedList[T any] struct {
	first *item[T]
	last  *item[T]
	size  int
}

func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) Add(v T) {
	curr := &item[T]{v: v}
	if l.last == nil {
		l.first = curr
		l.last = curr
	} else {
		curr.prev = l.last
		l.last.next = curr
		l.last = curr
	}
	l.size++
}

func (l *DoublyLinkedList[T]) Get(idx int) (T, bool) {
	var zero T
	if l.first == nil || idx < 0 || idx >= l.size {
		return zero, false
	}
	curr := l.first
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	return curr.v, true
}

func (l *DoublyLinkedList[T]) Remove(idx int) {
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

func (l *DoublyLinkedList[T]) Values() []T {
	res := make([]T, l.size)
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