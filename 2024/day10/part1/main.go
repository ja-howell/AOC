package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// data := getInput("smallInput.txt")
	data := getInput("input.txt")
	grid := convertToGrid(data)
	// fmt.Println(grid)
	fmt.Println(getTrailScore(grid))
}

type Direction struct {
	r int
	c int
}

var (
	left  = Direction{c: -1, r: 0}
	right = Direction{c: 1, r: 0}
	up    = Direction{c: 0, r: -1}
	down  = Direction{c: 0, r: 1}
)

func getInput(file string) [][]byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	grid := [][]byte{}

	for s.Scan() {
		grid = append(grid, []byte(s.Text()))
	}
	return grid
}

func getTrailScore(grid [][]int) int {
	score := 0
	for r, row := range grid {
		for c, col := range row {
			if col == 0 {
				score += findTrails(r, c, grid, map[Direction]struct{}{})
			}
		}
	}
	return score
}

func findTrails(r int, c int, grid [][]int, summits map[Direction]struct{}) (score int) {
	//we've got a trailhead at grid[r][c]
	score = 0
	if grid[r][c] == 9 {
		pos := Direction{r: r, c: c}
		_, ok := summits[pos]
		if ok {
			return 0
		}
		summits[pos] = struct{}{}
		return 1
	}
	space := grid[r][c]

	//check left
	if isInBounds(r, c+left.c, grid) && grid[r][c+left.c] == space+1 {
		score += findTrails(r, c+left.c, grid, summits)
	}
	//check right
	if isInBounds(r, c+right.c, grid) && grid[r][c+right.c] == space+1 {
		score += findTrails(r, c+right.c, grid, summits)
	}
	//check up
	if isInBounds(r+up.r, c, grid) && grid[r+up.r][c] == space+1 {
		score += findTrails(r+up.r, c, grid, summits)
	}
	//check down
	if isInBounds(r+down.r, c, grid) && grid[r+down.r][c] == space+1 {
		score += findTrails(r+down.r, c, grid, summits)
	}

	return score
}

func convertToGrid(data [][]byte) [][]int {
	grid := [][]int{}
	for _, row := range data {
		grid = append(grid, convertRow(row))
	}
	return grid
}

func convertRow(data []byte) []int {
	row := []int{}
	for _, val := range data {
		row = append(row, convertBytetoInt(val))
	}
	return row
}

func convertBytetoInt(x byte) int {
	return int(x & 0xf)
}

func isInBounds(r int, c int, grid [][]int) bool {
	return r < len(grid) && r >= 0 && c < len(grid[r]) && c >= 0
}
