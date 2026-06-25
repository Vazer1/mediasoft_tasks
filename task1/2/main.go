package main

import "fmt"

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func main() {
	const size = 512
	empls := [size]*Employee{}
	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scan(&cmd)

		switch cmd {
		case 1:
			// Добавляем нового сотрудника
			empl := new(Employee)
			fmt.Println("\nИмя:")
			fmt.Scan(&empl.Name)
			fmt.Println("Возраст:")
			fmt.Scan(&empl.Age)
			fmt.Println("Позиция:")
			fmt.Scan(&empl.Position)
			fmt.Println("Зарплата:")
			fmt.Scan(&empl.Salary)
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					break
				}
			}
		case 2:
			deleteNum := 0
			fmt.Println("Введите индекс сотрудника для удаления")
			fmt.Scan(&deleteNum)
			for deleteNum>=size || deleteNum<0 {
				fmt.Println("Введите индекс сотрудника для удаления, меньше 512")
				fmt.Scan(&deleteNum)
			}
			
			if empls[deleteNum]!=nil{
				empls[deleteNum] = nil
			}
		case 3:
			for i := 0; i < size; i++ {
				if empls[i] != nil {
					fmt.Print(i)
					fmt.Println(*empls[i])
				}
			}
		case 4:
			return
		}
	}
}
