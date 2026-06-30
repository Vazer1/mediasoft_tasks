package main

import (
	"fmt"

	"github.com/Vazer1/mediasoft_tasks/task3/list/list"
)

var commands = `
1 - add
2 - get
3 - remove
4 - values
5 - Выйти из программы
Ваш выбор: `

func main() {
	myList := list.New[int]()

	for {
		var val, idx int
		cmd := 0
		fmt.Print(commands)
		fmt.Scanln(&cmd)

		switch cmd {
		case 1:
			fmt.Println("Введите число:")
			fmt.Scanln(&val)
			myList.Add(val)
		case 2:
			fmt.Println("Введите индекс:")
			fmt.Scanln(&idx)
			if res, ok := myList.Get(idx); ok {
				fmt.Println("Значение:", res)
			} else {
				fmt.Println("Индекс выходит за границы списка")
			}
		case 3:
			fmt.Println("Введите индекс:")
			fmt.Scanln(&idx)
			myList.Remove(idx)
		case 4:
			slice := myList.Values()
			fmt.Println("Текущий список:", slice)
		case 5:
			return
		default:
			fmt.Println("Неизвестная команда. Попробуйте еще раз.")
		}
	}
}