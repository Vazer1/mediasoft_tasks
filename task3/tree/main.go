package main

import (
	"fmt"

	"github.com/Vazer1/mediasoft_tasks/task3/tree/tree"
)

var commands = `
1 - add
2 - remove
3 - values
4 - Выйти из программы
Ваш выбор: `

func main() {
	myTree := tree.New[int]()

	for {
		var val int
		cmd := 0
		fmt.Print(commands)
		fmt.Scanln(&cmd)

		switch cmd {
		case 1:
			fmt.Println("Введите число:")
			fmt.Scanln(&val)
			myTree.Add(val) 
		case 2:
			fmt.Println("Введите число:")
			fmt.Scanln(&val)
			myTree.Remove(val)
		case 3:
			slice := myTree.Values()
			fmt.Println("Текущие значения:", slice)
		case 4:
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}