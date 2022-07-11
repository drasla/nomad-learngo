package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total = total + number
	}

	total = 0
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
		total = total + numbers[i]
	}

	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
