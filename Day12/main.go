package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	ActionNorth   byte = 'N'
	ActionSouth        = 'S'
	ActionEast         = 'E'
	ActionWest         = 'W'
	ActionLeft         = 'L'
	ActionRight        = 'R'
	ActionForward      = 'F'
)

type Instruction struct {
	Action byte
	Value  int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func parseInput(input string) []Instruction {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	instructions := make([]Instruction, len(lines))

	for i, line := range lines {
		val, _ := strconv.Atoi(line[1:])
		instructions[i] = Instruction{
			Action: line[0],
			Value:  val,
		}
	}

	return instructions
}

func solvePart1(instructions []Instruction) int {
	currentDir := 0
	position := []int{0, 0}
	offsets := map[byte][]int{
		ActionNorth: {-1, 0},
		ActionSouth: {+1, 0},
		ActionEast: {0, +1},
		ActionWest: {0, -1},
	}

	for _, instruction := range instructions {
		switch instruction.Action {
		case ActionNorth, ActionSouth, ActionEast, ActionWest:
			position[0] += offsets[instruction.Action][0] * instruction.Value
			position[1] += offsets[instruction.Action][1] * instruction.Value
		case ActionLeft:
			currentDir -= instruction.Value
			if currentDir < 0 {
				currentDir = 360 + currentDir
			}
		case ActionRight:
			currentDir = (currentDir + instruction.Value) % 360
		case ActionForward:
			switch currentDir {
			case 0:
				position[1] += instruction.Value
			case 90:
				position[0] += instruction.Value
			case 180:
				position[1] -= instruction.Value
			case 270:
				position[0] -= instruction.Value
			}
		}

		fmt.Printf("%c %d %v\n", instruction.Action, instruction.Value, position)
	}

	return abs(position[0]) + abs(position[1])
}

func solvePart2(instructions []Instruction) int {
	return -1
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	instructions := parseInput(string(input))

	fmt.Println(solvePart1(instructions))
	fmt.Println(solvePart2(instructions))
}
