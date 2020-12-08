package main

import "github.com/pkg/errors"

type MachineState struct {
	InstructionPointer int
	Accumulator int
}

func run(program []Operation) (MachineState, error) {
	seen := make(map[int]bool)
	machine := &MachineState{
		InstructionPointer: 0,
		Accumulator:        0,
	}

	l := len(program)
	for machine.InstructionPointer < l {
		if _, exists := seen[machine.InstructionPointer]; exists {
			return *machine, errors.New("Infinite Loop")
		}

		seen[machine.InstructionPointer] = true
		machine = program[machine.InstructionPointer].execute(machine)
	}

	return *machine, nil
}