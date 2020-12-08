package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func solvePart1(program []Operation) int {
	state, _ := run(program)
	return state.Accumulator
}

func solvePart2(program []Operation) int {
	for addr := range program {
		old := program[addr]
		switch old.(type) {
		case InsNop:
			program[addr] = InsJmp{
				Instruction: Instruction{
					Operation: OpJmp,
					Operand:   old.(InsNop).Operand,
				},
			}
		case InsJmp:
			program[addr] = InsNop{
				Instruction: Instruction{
					Operation: OpNop,
					Operand:   old.(InsJmp).Operand,
				},
			}
		default:
			continue
		}

		state, err := run(program)
		if err == nil {
			return state.Accumulator
		}

		// reset to previous state
		program[addr] = old
	}

	return -1
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	program := parseProgram(string(input))

	fmt.Println(solvePart1(program))
	fmt.Println(solvePart2(program))
}
