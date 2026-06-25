package main

import "fmt"

func main() {
	nums := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	fmt.Println("Введите римское число: ")
	var roman string
	fmt.Scanln(&roman)
	res := 0
	length := len(roman)

	for i := 0; i < length; i++ {
		currentVal := nums[rune(roman[i])]

		if i+1 < length && currentVal < nums[rune(roman[i+1])] {
			res -= currentVal
		} else {
			res += currentVal
		}
	}

	fmt.Println(res)
}