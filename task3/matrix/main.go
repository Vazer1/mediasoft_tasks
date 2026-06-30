package main

import (
	"math/rand"

	"github.com/Vazer1/mediasoft_tasks/task3/matrix/matrix"
)

func main() {
	rows, cols := 3, 3
	
	mat := matrix.New[int](rows, cols)
	numMap := make(map[int]bool)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for {
				num := rand.Intn(100)

				if !numMap[num] {
					numMap[num] = true
					
					mat.Set(i, j, num)
					break
				}
			}
		}
	}

	mat.Print()
}