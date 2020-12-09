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
	lines := strings.Split(strings.TrimSpace(input),"\n")
	numbers := make([]int, len(lines))

	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
	}

	return numbers
}

const PreambleSize = 25

func solvePart1(numbers []int) int {
	for i := PreambleSize; i < len(numbers); i++ {
		if !checkPreamble(numbers[i-PreambleSize:i], numbers[i]) {
			return numbers[i]
		}
	}

	return -1
}

func checkPreamble(preamble []int, current int) bool {
	for i := 0; i < PreambleSize; i++ {
		for j := i+1; j < PreambleSize; j++ {
			if preamble[i] + preamble[j] == current {
				return true
			}
		}
	}

	return false
}

func solvePart2(numbers []int, target int) int {
	for i := 0; i < len(numbers); i++ {
		for j := i+1; j < len(numbers); j++ {
			set := numbers[i:j+1]
			if sum := getSum(set); sum == target {
				sort.Ints(set)
				return set[0] + set[len(set) - 1]
			}
		}
	}

	return -1
}

func getSum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	numbers := parseInput(string(input))

	part1 := solvePart1(numbers)
	fmt.Println(part1)
	fmt.Println(solvePart2(numbers, part1))
}

