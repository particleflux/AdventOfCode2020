package main

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

const (
	EmptySeat    byte = 'L'
	OccupiedSeat byte = '#'
	Floor        byte = '.'
)

func solvePart1(grid [][]byte) int {
	width, height := len(grid[0]), len(grid)

	// use double buffering instead of continuously allocating
	tmp := make([][]byte, height)
	for j := range grid {
		tmp[j] = make([]byte, width)
		copy(tmp[j], grid[j])
	}

	current := &grid
	target := &tmp
	changed := true
	swp := current

	i := 0
	for ; changed; i++ {
		fmt.Println(string(bytes.Join(*current, []byte{0xa})))
		fmt.Println()
		changed = simulateTurn(current, target, width, height)
		swp = current
		current = target
		target = swp
	}

	fmt.Println("no change in round", i)
	numOccupied := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if b, _ := isState(current, x, y, width, height, OccupiedSeat); b {
				numOccupied++
			}
		}
	}

	return numOccupied
}

func simulateTurn(current *[][]byte, target *[][]byte, width, height int) bool {
	changed := false

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			neighbours := countOccupied(current, x, y, width, height)
			(*target)[y][x] = (*current)[y][x]

			switch (*current)[y][x] {
			case EmptySeat:
				// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
				if neighbours == 0 {
					changed = true
					(*target)[y][x] = OccupiedSeat
				}
			case OccupiedSeat:
				// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
				if neighbours >= 5 {
					changed = true
					(*target)[y][x] = EmptySeat
				}
			}
		}
	}

	return changed
}

func countOccupied(grid *[][]byte, x, y , width, height int) int {
	neighbours := [][]int{
		{1, 0},   // right
		{0, 1},   // bottom
		{-1, 0},  // left
		{0, -1},  // top
		{1, 1},   // bottom right
		{-1, 1},  // bottom left
		{-1, -1}, // upper left
		{1, -1},  // upper right
	}
	sum := 0
	for _, neighbour := range neighbours {
		var err error
		var isFloor bool

		xx, yy := x, y
		for err == nil {
			xx += neighbour[0]
			yy += neighbour[1]
			isFloor, err = isState(grid, xx, yy, width, height, Floor)
			if err == nil && !isFloor {
				if occupied, _ := isState(grid, xx, yy, width, height, OccupiedSeat); occupied {
					sum++
				}
				break
			}
		}
	}

	return sum
}

func isState(grid *[][]byte, x, y , width, height int, state byte) (bool, error) {
	if x < 0 || y < 0 || x >= width || y >= height {
		return false, errors.New("out of bounds")
	}

	return (*grid)[y][x] == state, nil
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	grid := bytes.Split(bytes.TrimSpace(input), []byte{0xa})
	fmt.Println(grid)

	fmt.Println(solvePart1(grid))
}
