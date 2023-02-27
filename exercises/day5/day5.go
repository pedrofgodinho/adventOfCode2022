package day5

import (
	"fmt"
	"strings"
)

type Solver struct{}

func parseStack(stack string) [][]rune {
	var res [][]rune

	lines := strings.Split(stack, "\n")
	for _, line := range lines[:len(lines)-1] {
		for i := 0; i*4+2 < len(line); i++ {
			if line[i*4] == '[' {
				char := line[i*4+1]
				for i >= len(res) {
					res = append(res, []rune{})
				}
				res[i] = append(res[i], rune(char))
			}
		}
	}
	for _, stack := range res {
		for i := 0; i < len(stack)/2; i++ {
			j := len(stack) - i - 1
			stack[i], stack[j] = stack[j], stack[i]
		}
	}

	return res
}

func printStack(stack [][]rune) {
	for i, line := range stack {
		fmt.Print(i+1, " ")
		for _, box := range line {
			fmt.Print(string(box))
		}
		fmt.Println()
	}
}

func makeMove(stack [][]rune, from, to int) {
	if from >= len(stack) || len(stack[from]) == 0 {
		return
	}
	stack[to] = append(stack[to], stack[from][len(stack[from])-1])
	stack[from] = stack[from][:len(stack[from])-1]
}

func makeBulkMove(stack [][]rune, amount, from, to int) {
	if from >= len(stack) || len(stack[from]) == 0 {
		return
	}
	for i := amount; i > 0; i-- {
		stack[to] = append(stack[to], stack[from][len(stack[from])-i])
	}
	stack[from] = stack[from][:len(stack[from])-amount]
}

func (s Solver) SolvePart1(input string) string {
	split := strings.Split(input, "\n\n")
	stackStr := split[0]
	proc := split[1]
	stack := parseStack(stackStr)
	printStack(stack)
	for _, move := range strings.Split(proc, "\n") {
		var count, from, to int
		fmt.Sscanf(move, "move %d from %d to %d", &count, &from, &to)
		from--
		to--
		for i := 0; i < count; i++ {
			makeMove(stack, from, to)
		}
	}
	fmt.Println()
	printStack(stack)
	var result string
	for _, box := range stack {
		result += string(box[len(box)-1])
	}
	return result
}

func (s Solver) SolvePart2(input string) string {
	split := strings.Split(input, "\n\n")
	stackStr := split[0]
	proc := split[1]
	stack := parseStack(stackStr)
	printStack(stack)
	for _, move := range strings.Split(proc, "\n") {
		var count, from, to int
		fmt.Sscanf(move, "move %d from %d to %d", &count, &from, &to)
		from--
		to--
		makeBulkMove(stack, count, from, to)
	}
	fmt.Println()
	printStack(stack)
	var result string
	for _, box := range stack {
		result += string(box[len(box)-1])
	}
	return result
}
