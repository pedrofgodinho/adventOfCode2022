package day4

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (s Solver) SolvePart1(input string) string {
	in := strings.Split(input, "\n")
	var counter int
	for _, line := range in {
		var a1, a2, b1, b2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a1, &a2, &b1, &b2)
		if (a1 <= b1 && a2 >= b2) || (b1 <= a1 && b2 >= a2) {
			counter += 1
		}
	}
	return strconv.Itoa(counter)
}

func (s Solver) SolvePart2(input string) string {
	in := strings.Split(input, "\n")
	var counter int
	for _, line := range in {
		var a1, a2, b1, b2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a1, &a2, &b1, &b2)
		if (a1 <= b1 && a2 >= b2) ||
			(b1 <= a1 && b2 >= a2) ||
			(a1 <= b2 && a2 >= b2) ||
			(a1 <= b1 && a2 >= b1) {
			counter += 1
		}
	}
	return strconv.Itoa(counter)
}
