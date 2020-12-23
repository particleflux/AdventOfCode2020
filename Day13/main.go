package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	target, _ := strconv.Atoi(lines[0])
	ids := strings.Split(lines[1], ",")
	numbers := make([]int, 0)

	for _, id := range ids {
		if id == "x" {
			continue
		}
		n, _ := strconv.Atoi(id)
		numbers = append(numbers, n)
	}

	return numbers, target
}

func solvePart1(ids []int, target int) int {
	best, bestDistance := -1, 9999

	for _, id := range ids {
		dist := ((target/id)+1)*id - target
		if dist < bestDistance {
			bestDistance = dist
			best = id
		}
	}

	return best * bestDistance
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	ids, target := parseInput(string(input))

	fmt.Println(solvePart1(ids, target))
}
