package day10

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

const (
	Noop = iota
	Addx
)

type Instruction struct {
	opcode   int
	argument int
}

type Cpu struct {
	x                int
	instructions     []Instruction
	ip               int
	cycle            int
	instructionCycle int
}

func New(instructions []Instruction) Cpu {
	return Cpu{x: 1, instructions: instructions}
}

func (c *Cpu) Step() bool {
	instruction := c.instructions[c.ip]
	switch instruction.opcode {
	case Noop:
		c.instructionCycle = 0
		c.cycle++
		c.ip++
	case Addx:
		if c.instructionCycle == 0 {
			c.instructionCycle++
			c.cycle++
		} else {
			c.instructionCycle = 0
			c.cycle++
			c.x += instruction.argument
			c.ip++
		}
	}
	return c.ip < len(c.instructions)
}

func (s Solver) SolvePart1(input string) string {
	var instructions []Instruction
	for _, str := range strings.Split(input, "\n") {
		var instruction string
		var arg int
		fmt.Sscanf(str, "%s %d", &instruction, &arg)

		var opcode int
		switch instruction {
		case "noop":
			opcode = Noop
		case "addx":
			opcode = Addx
		default:
			panic("Invalid Instruction!")
		}
		instructions = append(instructions, Instruction{opcode, arg})
	}
	cpu := New(instructions)

	var sum int
	for cpu.Step() {
		if ((cpu.cycle+1)-20)%40 == 0 {
			sum += (cpu.cycle + 1) * cpu.x
		}
	}
	return strconv.Itoa(sum)
}

func (s Solver) SolvePart2(input string) string {
	var instructions []Instruction
	for _, str := range strings.Split(input, "\n") {
		var instruction string
		var arg int
		fmt.Sscanf(str, "%s %d", &instruction, &arg)

		var opcode int
		switch instruction {
		case "noop":
			opcode = Noop
		case "addx":
			opcode = Addx
		default:
			panic("Invalid Instruction!")
		}
		instructions = append(instructions, Instruction{opcode, arg})
	}
	cpu := New(instructions)

	for next := true; next; next = cpu.Step() {
		crtCycle := cpu.cycle % 240
		diff := cpu.x - crtCycle%40
		if diff <= 1 && diff >= -1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
		if (crtCycle+1)%40 == 0 {
			fmt.Println()
		}
	}
	return ""
}
