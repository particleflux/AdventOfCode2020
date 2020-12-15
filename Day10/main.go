package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	numbers := make([]int, len(lines))

	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
	}

	return numbers
}

func solvePart1(numbers []int) int {
	diffs := []int{0, 0, 1}

	diffs[numbers[0]-1]++
	for i := 1; i < len(numbers); i++ {
		diffs[numbers[i]-numbers[i-1]-1]++
	}

	return diffs[0] * diffs[2]
}

func solvePart2(numbers []int) int {
	num := map[int]int{0: 1}
	var cur int

	for _, cur = range numbers {
		num[cur] = num[cur-1] + num[cur-2] + num[cur-3]
	}

	return num[cur]
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	numbers := parseInput(string(input))
	sort.Ints(numbers)

	fmt.Println(solvePart1(numbers))
	fmt.Println(solvePart2(numbers))
}
