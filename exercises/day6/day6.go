package day6

import (
	"fmt"
	"strconv"
)

type Solver struct{}

func (s Solver) SolvePart1(input string) string {
	var a, b, c, d rune
	for i, char := range input {
		a = b
		b = c
		c = d
		d = char
		if i > 2 &&
			a != b &&
			a != c &&
			a != d &&
			b != c &&
			b != d &&
			c != d {
			fmt.Printf("%c %c %c %c\n", a, b, c, d)
			return strconv.Itoa(i + 1)
		}
	}
	return "???"
}

func (s Solver) SolvePart2(input string) string {
	var chars [14]rune
	for i, char := range input {
		for i := 0; i < len(chars)-1; i++ {
			chars[i] = chars[i+1]
		}
		chars[len(chars)-1] = char
		if i < 12 {
			continue
		}
		hasDuplicates := false
		for j := 0; j < len(chars); j++ {
			for k := j + 1; k < len(chars); k++ {
				if chars[j] == chars[k] {
					hasDuplicates = true
				}
			}
		}
		if !hasDuplicates {
			return strconv.Itoa(i + 1)
		}
	}
	return "???"
}
