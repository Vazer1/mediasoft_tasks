package main

import (
	"fmt"

	"github.com/Vazer1/mediasoft_tasks/task3/roman/roman"
)

func main() {
	nums := roman.New[int]()

	fmt.Print("Введите римское число: ")
	var romanStr string
	fmt.Scanln(&romanStr)

	res := nums.Convert(romanStr)

	fmt.Println(res)
}