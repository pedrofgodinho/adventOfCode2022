package day2

import (
	"strconv"
	"strings"
)

type Solver struct{}

func (s Solver) SolvePart1(input string) string {
	var score int
	for _, line := range strings.Split(input, "\n") {
		me := int(line[2] - 'X')
		other := int(line[0] - 'A')
		var resultPoints int
		switch {
		case me == other:
			resultPoints += 3
		case (me == 1 && other == 0) || (me == 0 && other == 2) || (me == 2 && other == 1):
			resultPoints += 6
		default:
			resultPoints += 0
		}
		score += resultPoints
		score += me + 1
	}
	return strconv.Itoa(score)
}

func (s Solver) SolvePart2(input string) string {
	var score int
	scores := map[string]int{"B X": 1, "C X": 2, "A X": 3, "A Y": 4, "B Y": 5, "C Y": 6, "C Z": 7, "A Z": 8, "B Z": 9}
	for _, line := range strings.Split(input, "\n") {
		score += scores[line]
	}
	return strconv.Itoa(score)
}
