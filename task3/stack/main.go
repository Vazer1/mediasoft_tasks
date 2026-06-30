package main

import (
	"fmt"

	"github.com/Vazer1/mediasoft_tasks/task3/stack/stack"
)

var commands = `
1 - push
2 - pop
3 - peek
4 - Выйти из программы
Ваш выбор: `

func main() {
	myStack := stack.New[int](256)

	for {
		var val int
		cmd := 0
		fmt.Print(commands)
		fmt.Scanln(&cmd)

		switch cmd {
		case 1:
			fmt.Println("Введите число:")
			fmt.Scanln(&val)
			myStack.Push(val) 
			
		case 2:
			val, ok := myStack.Pop()
			if ok {
				fmt.Println("Извлечено:", val)
			} else {
				fmt.Println("Ошибка: стек пуст!")
			}
			
		case 3:
			val, ok := myStack.Peek()
			if ok {
				fmt.Println("На вершине:", val)
			} else {
				fmt.Println("Ошибка: стек пуст!")
			}
			
		case 4:
			return
			
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}