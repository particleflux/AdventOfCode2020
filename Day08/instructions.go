package main

import "fmt"

type Opcode string

const (
	OpAcc Opcode = "acc"
	OpJmp Opcode = "jmp"
	OpNop Opcode = "nop"
)

type Operation interface {
	execute(state *MachineState) *MachineState
}

type Instruction struct {
	Operation Opcode
	Operand int
}

func (i Instruction) string() string {
	return fmt.Sprintf("%s %d", i.Operation, i.Operand)
}

type InsNop struct {
	Instruction
}

func (i InsNop) execute(state *MachineState) *MachineState {
	state.InstructionPointer++
	return state
}

type InsAcc struct {
	Instruction
}

func (i InsAcc) execute(state *MachineState) *MachineState {
	state.InstructionPointer++
	state.Accumulator += i.Operand
	return state
}

type InsJmp struct {
	Instruction
}

func (i InsJmp) execute(state *MachineState) *MachineState {
	state.InstructionPointer += i.Operand
	return state
}