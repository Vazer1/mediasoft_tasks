package main

import "fmt"

type stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int   // индекс головы стека
}

func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

// push - добавление в стек значения
func push(s *stack, v any) {
	if s.head >= len(s.s)-1 {
		return
	}
	s.head++
	s.s[s.head] = v
}

// pop - получения значения из стека и его удаление из вершины
func pop(s *stack) any {
	if s.head == -1 {
		return nil
	}
	res := s.s[s.head]
	s.s[s.head] = nil
	s.head--

	return res
}

// peek - просмотр значения на вершине стека
func peek(s *stack) any {
	if s.head == -1 {
		return nil
	}

	return s.s[s.head]
}

var commands = `
1 - push
2 - pop
3 - peek
4 - Выйти из программы
`
func main() {
	myStack := newStack(256)

	for {
		var val int
		cmd := 0
		fmt.Print(commands)
		fmt.Scanln(&cmd)

		switch cmd {
		case 1:
			fmt.Println("Введите число")
			fmt.Scanln(&val)
			push(myStack, val)
		case 2:
			val := pop(myStack)
			fmt.Println(val)
		case 3:
			val := peek(myStack)
			fmt.Println(val)
		case 4:
			return
		}
	}
}