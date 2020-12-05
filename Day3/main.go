package main

import (
	"fmt"
)

const (
	SymbolTree  byte = '#'
	SymbolEmpty      = '.'
)

func checkPath(grid [][]byte, height, width int, path []int) int {
	numTrees := 0
	for x, y := 0, 0; y < height; {
		if grid[y][x] == SymbolTree {
			numTrees++
		}

		x = (x + path[0]) % width
		y += path[1]
	}

	return numTrees
}

func main() {
	grid := make([][]byte, 400)
	height := 0

	for {
		if _, err := fmt.Scanln(&grid[height]); err != nil {
			break
		}
		height++
	}

	width := len(grid[0])
	fmt.Println("part 1:", checkPath(grid, height, width, []int{3, 1}))

	part2 := checkPath(grid, height, width, []int{1, 1})
	part2 *= checkPath(grid, height, width, []int{3, 1})
	part2 *= checkPath(grid, height, width, []int{5, 1})
	part2 *= checkPath(grid, height, width, []int{7, 1})
	part2 *= checkPath(grid, height, width, []int{1, 2})
	fmt.Println("part 2:", part2)
}
