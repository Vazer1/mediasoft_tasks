package stack

type Stack[T any] struct {
	s    []T
	head int
}

func New[T any](size int) *Stack[T] {
	return &Stack[T]{
		s:    make([]T, size),
		head: -1,
	}
}

func (st *Stack[T]) Push(v T) {
	if st.head >= len(st.s)-1 {
		return // Стек переполнен
	}
	st.head++
	st.s[st.head] = v
}

func (st *Stack[T]) Pop() (T, bool) {
	var zero T

	if st.head == -1 {
		return zero, false
	}

	res := st.s[st.head]
	st.s[st.head] = zero
	st.head--

	return res, true
}

func (st *Stack[T]) Peek() (T, bool) {
	var zero T

	if st.head == -1 {
		return zero, false
	}

	return st.s[st.head], true
}