package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Direction struct {
	x int
	y int
}

type Position struct {
	x int
	y int
}

type State struct {
	facing string
	pos    Position
}

var (
	left  = Direction{x: -1, y: 0}
	right = Direction{x: 1, y: 0}
	up    = Direction{x: 0, y: -1}
	down  = Direction{x: 0, y: 1}
)

func main() {
	grid := getInput("input.txt")
	// printGrid(grid)

	total := countValidObstacles(grid)

	fmt.Println(total)

}

func getInput(file string) [][]byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	grid := [][]byte{}

	for s.Scan() {
		grid = append(grid, []byte(s.Text()))
	}

	return grid
}

func getVisitedPositions(grid [][]byte, guard State) []Position {
	visited := []Position{guard.pos}

	for isInBounds(guard.pos.x, len(grid[0]), guard.pos.y, len(grid)) {
		if !slices.Contains(visited, guard.pos) {
			visited = append(visited, guard.pos)
		}
		guard = moveGuard(guard, grid)
	}
	return visited
}

func countValidObstacles(grid [][]byte) int {
	guard := getStartingState(grid)
	visited := getVisitedPositions(grid, guard)
	//check each visited position
	//change each positon to a '#'
	//change it back to a '.'
	total := 0
	for _, pos := range visited {
		grid[pos.y][pos.x] = '#'
		if containsCycle(grid, guard) {
			total++
		}
		grid[pos.y][pos.x] = '.'
	}
	return total
}

func getStartingState(grid [][]byte) State {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' {
				return State{facing: "n", pos: Position{x: j, y: i}}
			}
		}
	}
	return State{facing: "n", pos: Position{x: 0, y: 0}}
}

func isInBounds(x int, xLength int, y int, yLength int) bool {
	if x < 0 || x >= xLength {
		return false
	} else if y < 0 || y >= yLength {
		return false
	}

	return true
}

func moveGuard(guard State, grid [][]byte) State {
	//check for obstacle
	//if obstacle ahead turn 90 degrees
	// check for obstacle again
	// repeat until no obstacle
	//move forward in direction facing
	nextPos := getNextPosistion(guard)
	if isInBounds(nextPos.x, len(grid[0]), nextPos.y, len(grid)) {
		for grid[nextPos.y][nextPos.x] == '#' {
			guard.facing = rotate90Degrees(guard)
			nextPos = getNextPosistion(guard)
		}
	}
	guard.pos.x = nextPos.x
	guard.pos.y = nextPos.y

	return guard

}

func getNextPosistion(guard State) Position {
	currPos := guard.pos
	if guard.facing == "n" {
		return Position{x: currPos.x + up.x, y: currPos.y + up.y}
	} else if guard.facing == "s" {
		return Position{x: currPos.x + down.x, y: currPos.y + down.y}
	} else if guard.facing == "e" {
		return Position{x: currPos.x + right.x, y: currPos.y + right.y}
	} else {
		return Position{x: currPos.x + left.x, y: currPos.y + left.y}
	}
}

func rotate90Degrees(guard State) string {
	if guard.facing == "n" {
		return "e"
	} else if guard.facing == "s" {
		return "w"
	} else if guard.facing == "e" {
		return "s"
	} else {
		return "n"
	}
}

func containsCycle(grid [][]byte, guard State) bool {
	statesSet := map[State]struct{}{}
	for isInBounds(guard.pos.x, len(grid[0]), guard.pos.y, len(grid)) {
		_, ok := statesSet[guard]
		if ok {
			return true
		}
		statesSet[guard] = struct{}{}
		guard = moveGuard(guard, grid)
	}
	return false
}

// func printGrid(grid [][]byte) {
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[i]); j++ {
// 			fmt.Print(string(grid[i][j]))
// 		}
// 		fmt.Println()
// 	}
// }
