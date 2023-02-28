package main

import (
	"fmt"
	"github.com/pedrofgodinho/adventOfCode2022/exercises"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day1"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day10"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day11"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day12"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day2"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day3"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day4"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day5"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day6"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day7"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day8"
	"github.com/pedrofgodinho/adventOfCode2022/exercises/day9"
	"os"
)

func main() {
	var day, part string
	fmt.Print("Day: ")
	_, err := fmt.Scanln(&day)
	if err != nil {
		return
	}
	fmt.Print("Part: ")
	_, err = fmt.Scanln(&part)
	if err != nil {
		return
	}
	inputFileName := "inputs/day" + day

	fileContent, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	solvers := map[string]exercises.Exercise{
		"1":  day1.Solver{},
		"2":  day2.Solver{},
		"3":  day3.Solver{},
		"4":  day4.Solver{},
		"5":  day5.Solver{},
		"6":  day6.Solver{},
		"7":  day7.Solver{},
		"8":  day8.Solver{},
		"9":  day9.Solver{},
		"10": day10.Solver{},
		"11": day11.Solver{},
		"12": day12.Solver{},
	}

	exercise := solvers[day]

	if part == "1" {
		fmt.Println(exercise.SolvePart1(string(fileContent)))
	} else if part == "2" {
		fmt.Println(exercise.SolvePart2(string(fileContent)))
	}
}
