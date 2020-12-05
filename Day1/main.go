package main

import (
	"fmt"
)

func solvePart1(numbers []int) int {
	l := len(numbers)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if numbers[i]+numbers[j] == 2020 {
				return numbers[i] * numbers[j]
			}
		}
	}

	return -1
}

func solvePart2(numbers []int) int {
	l := len(numbers)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if numbers[i] == 0 || numbers[j] == 0 || numbers[k] == 0 {
					continue
				}
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					return numbers[i] * numbers[j] * numbers[k]
				}
			}
		}
	}

	return -1
}

func main() {
	numbers := make([]int, 200)
	var num int

	for {
		if _, err := fmt.Scan(&num); err != nil {
			break
		}

		numbers = append(numbers, num)
	}

	fmt.Println(solvePart1(numbers))
	fmt.Println(solvePart2(numbers))
}
