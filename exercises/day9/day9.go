package day9

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

type Point struct {
	x int
	y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (p Point) Distance(other Point) int {
	diff := p.Sub(other)
	return max(abs(diff.x), abs(diff.y))
}

func (p Point) Magnitude() int {
	return abs(p.Distance(Point{0, 0}))
}

func (p Point) Add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

func (p Point) Sub(other Point) Point {
	return Point{p.x - other.x, p.y - other.y}
}

func (p Point) ClampToUnit() Point {
	res := Point{p.x, p.y}
	if res.x > 1 {
		res.x = 1
	} else if res.x < -1 {
		res.x = -1
	}
	if res.y > 1 {
		res.y = 1
	} else if res.y < -1 {
		res.y = -1
	}

	return res
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Snake struct {
	parts      []Point
	maxVisited Point
	minVisited Point
}

func New(size int) *Snake {
	parts := make([]Point, size)
	return &Snake{parts: parts}
}

func (s *Snake) Move(direction Point) {
	s.parts[0] = s.parts[0].Add(direction)
	for i := 1; i < len(s.parts); i++ {
		if s.parts[i-1].Distance(s.parts[i]) > 1 {
			s.parts[i] = s.parts[i].Add(s.parts[i-1].Sub(s.parts[i]).ClampToUnit())
		}
	}
}

func (s Solver) SolvePart1(input string) string {
	snake := New(2)
	tailVisits := make(map[Point]bool)
	tailVisits[snake.parts[0]] = true
	for _, move := range strings.Split(input, "\n") {
		var directionRune rune
		var amount int
		_, _ = fmt.Sscanf(move, "%c %d", &directionRune, &amount)
		var direction Point
		switch directionRune {
		case 'R':
			direction = Point{1, 0}
		case 'L':
			direction = Point{-1, 0}
		case 'U':
			direction = Point{0, 1}
		case 'D':
			direction = Point{0, -1}
		default:
			panic("Invalid Direction: " + string(directionRune))
		}
		for i := 0; i < amount; i++ {
			snake.Move(direction)
			tailVisits[snake.parts[len(snake.parts)-1]] = true
		}
	}
	return strconv.Itoa(len(tailVisits))
}

func (s Solver) SolvePart2(input string) string {
	snake := New(10)
	tailVisits := make(map[Point]bool)
	tailVisits[snake.parts[0]] = true
	for _, move := range strings.Split(input, "\n") {
		var directionRune rune
		var amount int
		_, _ = fmt.Sscanf(move, "%c %d", &directionRune, &amount)
		var direction Point
		switch directionRune {
		case 'R':
			direction = Point{1, 0}
		case 'L':
			direction = Point{-1, 0}
		case 'U':
			direction = Point{0, 1}
		case 'D':
			direction = Point{0, -1}
		default:
			panic("Invalid Direction: " + string(directionRune))
		}
		for i := 0; i < amount; i++ {
			snake.Move(direction)
			tailVisits[snake.parts[len(snake.parts)-1]] = true
		}
	}
	return strconv.Itoa(len(tailVisits))
}
