package day1

import (
	"strconv"
	"strings"
)

type Solver struct{}

func (s Solver) SolvePart1(input string) string {
	max := 0
	for _, elf := range strings.Split(input, "\n\n") {
		current := 0
		for _, meal := range strings.Split(elf, "\n") {
			cal, err := strconv.Atoi(meal)
			if err != nil {
				break
			}
			current += cal
		}
		if current >= max {
			max = current
		}
	}
	return strconv.Itoa(max)
}

func (s Solver) SolvePart2(input string) string {
	var max [3]int
	for _, elf := range strings.Split(input, "\n\n") {
		current := 0
		for _, meal := range strings.Split(elf, "\n") {
			cal, err := strconv.Atoi(meal)
			if err != nil {
				break
			}
			current += cal
		}
		if current > max[0] {
			max[1] = max[0]
			max[2] = max[1]
			max[0] = current
		} else if current > max[1] {
			max[2] = max[1]
			max[1] = current
		} else if current > max[2] {
			max[2] = current
		}
	}
	return strconv.Itoa(max[0] + max[1] + max[2])
}
