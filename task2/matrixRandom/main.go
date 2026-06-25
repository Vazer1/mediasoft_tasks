package main

import (
	"fmt"
	"math/rand"
)

func main() {

	matrix := [3][3]int{}
	numMap := make(map[int]bool)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for {
				num := rand.Intn(100)

				if _, ok := numMap[num]; !ok {
					numMap[num] = true 
					matrix[i][j] = num 
					break
				}
			}
		}
	}

	for _, row := range matrix {
		fmt.Println(row)
	}
}