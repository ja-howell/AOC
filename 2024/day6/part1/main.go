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

	total := countPositions(grid)

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

func countPositions(grid [][]byte) int {
	guard := getStartingState(grid)
	visited := []Position{guard.pos}
	fmt.Printf("Starting Position: x: %d y: %d facing: %s\n", guard.pos.x, guard.pos.y, guard.facing)

	//while guard is still on the map
	for isInBounds(guard.pos.x, len(grid[0]), guard.pos.y, len(grid)) {
		if !slices.Contains(visited, guard.pos) {
			visited = append(visited, guard.pos)
		}
		guard = moveGuard(guard, grid)
	}
	return len(visited)
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

// func printGrid(grid [][]byte) {
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[i]); j++ {
// 			fmt.Print(string(grid[i][j]))
// 		}
// 		fmt.Println()
// 	}
// }
