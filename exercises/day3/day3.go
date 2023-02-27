package day3

import (
	"strconv"
	"strings"
)

type Solver struct{}

func intersect(original string, target string) (rune, bool) {
	m := make(map[rune]bool)
	for _, item := range original {
		m[item] = true
	}
	for _, item := range target {
		if _, ok := m[item]; ok {
			return item, true
		}
	}
	return 0, false
}

func intersect3(original string, target1 string, target2 string) (rune, bool) {
	m1 := make(map[rune]bool)
	m2 := make(map[rune]bool)
	for _, item := range target1 {
		m1[item] = true
	}
	for _, item := range target2 {
		m2[item] = true
	}
	for _, item := range original {
		if _, ok := m1[item]; ok {
			if _, ok2 := m2[item]; ok2 {
				return item, true
			}
		}
	}
	return 0, false
}

func (s Solver) SolvePart1(input string) string {
	var prio int
	for _, line := range strings.Split(input, "\n") {
		line1 := line[:len(line)/2]
		line2 := line[len(line)/2:]
		item, _ := intersect(line1, line2)
		if item <= 'Z' {
			prio += int(item - 'A' + 27)
		} else {
			prio += int(item - 'a' + 1)
		}
	}
	return strconv.Itoa(prio)
}

func (s Solver) SolvePart2(input string) string {
	var prio int
	in := strings.Split(input, "\n")
	for i := 0; i < len(in); i += 3 {
		line1 := in[i]
		line2 := in[i+1]
		line3 := in[i+2]
		item, _ := intersect3(line1, line2, line3)
		if item <= 'Z' {
			prio += int(item - 'A' + 27)
		} else {
			prio += int(item - 'a' + 1)
		}
	}
	return strconv.Itoa(prio)
}
