package main

import "fmt"

func solvePart1(numbers []int, target int) int {
	lastOccurence := make(map[int]int, 0)

	// init with starting numbers
	for pos, number := range numbers {
		lastOccurence[number] = pos + 1
	}

	current := numbers[len(numbers) - 1]
	for i := len(numbers); i < target; i++ {
		next := 0
		if pos, exists := lastOccurence[current]; exists {
			next = i - pos
		}

		lastOccurence[current] = i
		current = next
	}

	return current
}

func main() {
	//fmt.Println(solvePart1([]int{0,3,6}))
	fmt.Println(solvePart1([]int{2,0,1,7,4,14,18}, 2020))
	fmt.Println(solvePart1([]int{2,0,1,7,4,14,18}, 30000000))
}
