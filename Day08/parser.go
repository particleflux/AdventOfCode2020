package main

import (
	"strconv"
	"strings"
)

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func parseInstruction(instruction string) Operation {
	parts := strings.Fields(instruction)

	inst := Instruction{
		Operation: Opcode(parts[0]),
		Operand:   MustAtoi(parts[1]),
	}

	switch Opcode(parts[0]) {
	case OpNop:
		return InsNop{inst}
	case OpAcc:
		return InsAcc{inst}
	case OpJmp:
		return InsJmp{inst}
	}

	panic("Unknown instruction " + instruction)
}

func parseProgram(program string) []Operation {
	operations := make([]Operation, 0)

	for _, instruction := range strings.Split(strings.TrimSpace(program), "\n") {
		operations = append(operations, parseInstruction(instruction))
	}

	return operations
}