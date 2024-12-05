package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction struct {
	x int
	y int
}

var (
	left      = Direction{x: -1, y: 0}
	right     = Direction{x: 1, y: 0}
	up        = Direction{x: 0, y: -1}
	down      = Direction{x: 0, y: 1}
	upLeft    = Direction{x: -1, y: -1}
	upRight   = Direction{x: 1, y: -1}
	downLeft  = Direction{x: -1, y: 1}
	downRight = Direction{x: 1, y: 1}
)

func getInput(file string) [][]byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	crossword := [][]byte{}
	for scanner.Scan() {
		crossword = append(crossword, []byte(scanner.Text()))
	}

	return crossword

}

func processCrossword(crossword [][]byte) int {
	total := 0
	for r := 0; r < len(crossword); r++ {
		for c := 0; c < len(crossword[r]); c++ {
			total += countXMASForCell(crossword, r, c)
		}
	}
	return total
}

func countXMASForCell(crossword [][]byte, r int, c int) int {
	total := 0
	directions := []Direction{
		left, right, up, down,
		upLeft, upRight, downLeft, downRight,
	}

	for _, dir := range directions {
		if isXmas(crossword, r, c, dir) {
			total++
		}
	}

	return total
}

func isXmas(crossword [][]byte, r int, c int, dir Direction) bool {
	xmas := []byte{'X', 'M', 'A', 'S'}
	for _, v := range xmas {
		//if out of bounds
		if r < 0 || r >= len(crossword) {
			return false
		}
		if c < 0 || c >= len(crossword) {
			return false
		}
		if v != crossword[r][c] {
			return false
		}
		r += dir.x
		c += dir.y
	}
	return true
}

func main() {
	crossword := getInput("input.txt")
	total := processCrossword(crossword)
	fmt.Println(total)

}
