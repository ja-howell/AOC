package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	grid := getInput("input.txt")
	// grid := getInput("smallInput.txt")
	antennas := findAntennas(grid)
	antinodes := findAntinodes(len(grid), len(grid[0]), antennas)
	fmt.Println(len(antinodes))
	// validAntinodes := validateAntinodes(len(grid), len(grid[0]), antinodes)

	// fmt.Println(len(validAntinodes))
}

type Cell struct {
	col int
	row int
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

func findAntennas(grid [][]byte) map[byte][]Cell {
	antennas := map[byte][]Cell{}

	for r, row := range grid {
		for c, val := range row {
			if val != '.' {
				antennas[val] = append(antennas[val], Cell{col: c, row: r})
			}
		}
	}

	return antennas
}

func findAntinodes(gl int, gw int, antennas map[byte][]Cell) map[Cell]struct{} {
	//should return a set of cells
	//takes a map of antennas
	antinodes := map[Cell]struct{}{}
	for _, antenna := range antennas {
		for i := 0; i < len(antenna); i++ {
			for j := i + 1; j < len(antenna); j++ {
				dc := antenna[i].col - antenna[j].col
				dr := antenna[i].row - antenna[j].row
				currCell := getAntinode(antenna[i], dr, dc)
				//while inbounds
				for isInbounds(gl, gw, currCell) {
					antinodes[currCell] = struct{}{}
					currCell = getAntinode(currCell, dr, dc)
				}
				currCell = getAntinode(antenna[j], -dr, -dc)
				for isInbounds(gl, gw, currCell) {
					antinodes[currCell] = struct{}{}
					currCell = getAntinode(currCell, -dr, -dc)
				}
			}
		}
	}
	return antinodes
}

func getAntinode(antenna Cell, dr int, dc int) Cell {
	return Cell{col: antenna.col + dc, row: antenna.row + dr}
}

func validateAntinodes(length int, width int, antinodes map[Cell]struct{}) map[Cell]struct{} {
	validAntinodes := map[Cell]struct{}{}
	for antinode := range antinodes {
		if validateAntinode(length, width, antinode) {
			validAntinodes[antinode] = struct{}{}
		}
	}
	return validAntinodes
}

func validateAntinode(length int, width int, antinode Cell) bool {
	return antinode.row >= 0 && antinode.row < width && antinode.col >= 0 && antinode.col < length
}

func isInbounds(length int, width int, antinode Cell) bool {
	return antinode.row >= 0 && antinode.row < width && antinode.col >= 0 && antinode.col < length
}
