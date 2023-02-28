package day12

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

type Point struct {
	row int
	col int
}

type Cell struct {
	steps     int
	direction byte
}

func bfs(maze [][]byte, start Point, end map[Point]bool) int {
	queue := list.New()
	queue.PushBack(start)

	explored := make(map[Point]bool)
	explored[start] = true

	solvedMaze := make([][]Cell, len(maze))
	for i := 0; i < len(maze); i++ {
		solvedMaze[i] = make([]Cell, len(maze[i]))
		for j := 0; j < len(maze[0]); j++ {
			solvedMaze[i][j] = Cell{direction: '.'}
		}
	}

	var found Point

	for queue.Len() > 0 {
		currentPoint := queue.Front().Value.(Point)
		queue.Remove(queue.Front())
		if _, ok := end[currentPoint]; ok {
			found = currentPoint
			break
		}
		offsets := map[Point]byte{
			{1, 0}:  '^',
			{-1, 0}: 'v',
			{0, 1}:  '<',
			{0, -1}: '>',
		}
		for offset, arrow := range offsets {
			nextPoint := Point{currentPoint.row + offset.row, currentPoint.col + offset.col}
			if _, nextExplored := explored[nextPoint]; !nextExplored &&
				nextPoint.row < len(maze) && nextPoint.row >= 0 &&
				nextPoint.col < len(maze[nextPoint.row]) && nextPoint.col >= 0 {
				if currentHeight, nextHeight := maze[currentPoint.row][currentPoint.col], maze[nextPoint.row][nextPoint.col]; currentHeight-1 <= nextHeight {
					explored[nextPoint] = true
					solvedMaze[nextPoint.row][nextPoint.col].direction = arrow
					solvedMaze[nextPoint.row][nextPoint.col].steps = solvedMaze[currentPoint.row][currentPoint.col].steps + 1
					queue.PushBack(nextPoint)
				}
			}
		}
	}

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[0]); j++ {
			fmt.Printf("%c", solvedMaze[i][j].direction)
		}
		fmt.Println()
	}
	return solvedMaze[found.row][found.col].steps
}

func (s Solver) SolvePart1(input string) string {
	split := strings.Split(input, "\n")
	maze := make([][]byte, len(split))

	start, end := Point{-1, -1}, Point{-1, -1}
	for row := 0; row < len(split); row++ {
		maze[row] = []byte(split[row])
		for col := 0; col < len(split[0]); col++ {
			if split[row][col] == 'S' {
				start.row = row
				start.col = col
			} else if split[row][col] == 'E' {
				end.row = row
				end.col = col
			}
		}
	}
	maze[start.row][start.col] = 'a'
	maze[end.row][end.col] = 'z'
	return strconv.Itoa(bfs(maze, end, map[Point]bool{start: true}))
}

func (s Solver) SolvePart2(input string) string {
	split := strings.Split(input, "\n")
	maze := make([][]byte, len(split))

	starts := make(map[Point]bool)
	end := Point{-1, -1}
	for row := 0; row < len(split); row++ {
		maze[row] = []byte(split[row])
		for col := 0; col < len(split[0]); col++ {
			if split[row][col] == 'S' {
				starts[Point{row, col}] = true
				maze[row][col] = 'a'
			} else if split[row][col] == 'a' {
				starts[Point{row, col}] = true
			} else if split[row][col] == 'E' {
				end.row = row
				end.col = col
			}
		}
	}
	maze[end.row][end.col] = 'z'

	return strconv.Itoa(bfs(maze, end, starts))
}
