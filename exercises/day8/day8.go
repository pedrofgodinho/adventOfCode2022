package day8

import (
	"strconv"
	"strings"
)

type Solver struct{}

// Very ugly, bruteforce-y solutions, but I didn't like this problem and can't be bothered to write a proper solution

func (s Solver) SolvePart1(input string) string {
	trees := strings.Split(input, "\n")

	var count int
	for row := 1; row < len(trees)-1; row++ {
		for col := 1; col < len(trees[row])-1; col++ {
			// left
			visible := true
			for check := col - 1; check >= 0; check-- {
				if trees[row][col] <= trees[row][check] {
					visible = false
					break
				}
			}
			if visible {
				count += 1
				continue
			}
			// right
			visible = true
			for check := col + 1; check < len(trees[row]); check++ {
				if trees[row][col] <= trees[row][check] {
					visible = false
					break
				}
			}
			if visible {
				count += 1
				continue
			}
			// up
			visible = true
			for check := row - 1; check >= 0; check-- {
				if trees[row][col] <= trees[check][col] {
					visible = false
					break
				}
			}
			if visible {
				count += 1
				continue
			}
			// down
			visible = true
			for check := row + 1; check < len(trees); check++ {
				if trees[row][col] <= trees[check][col] {
					visible = false
					break
				}
			}
			if visible {
				count += 1
				continue
			}
		}
	}
	return strconv.Itoa(count + len(trees)*2 + len(trees[0])*2 - 4)
}

func (s Solver) SolvePart2(input string) string {
	trees := strings.Split(input, "\n")

	var max int
	for row := 1; row < len(trees)-1; row++ {
		for col := 1; col < len(trees[row])-1; col++ {
			total := 1
			// left
			score := 0
			for check := col - 1; check >= 0; check-- {
				if trees[row][col] > trees[row][check] {
					score++
				} else {
					score++
					break
				}
			}
			total *= score
			// right
			score = 0
			for check := col + 1; check < len(trees[row]); check++ {
				if trees[row][col] > trees[row][check] {
					score++
				} else {
					score++
					break
				}
			}
			total *= score
			// up
			score = 0
			for check := row - 1; check >= 0; check-- {
				if trees[row][col] > trees[check][col] {
					score++
				} else {
					score++
					break
				}
			}
			total *= score

			// down
			score = 0
			for check := row + 1; check < len(trees); check++ {
				if trees[row][col] > trees[check][col] {
					score++
				} else {
					score++
					break
				}
			}
			total *= score
			if total > max {
				max = total
			}
		}
	}
	return strconv.Itoa(max)
}
