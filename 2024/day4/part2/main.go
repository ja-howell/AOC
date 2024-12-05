package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Direction struct {
	x int
	y int
}

var (
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
	for r := 1; r < len(crossword)-1; r++ {
		for c := 1; c < len(crossword[r])-1; c++ {
			if isCellXmas(crossword, r, c) {
				total++
			}
		}
	}
	return total
}

func isCellXmas(crossword [][]byte, r int, c int) bool {
	downRightDiagonal := getWord(crossword, r-1, c-1, downRight)
	downLeftDiagonal := getWord(crossword, r-1, c+1, downLeft)
	if downLeftDiagonal != "SAM" && downLeftDiagonal != "MAS" {
		return false
	}
	if downRightDiagonal != "SAM" && downRightDiagonal != "MAS" {
		return false
	}

	return true
}

func getWord(crossword [][]byte, r int, c int, dir Direction) string {
	sb := strings.Builder{}
	for i := 0; i < 3; i++ {
		sb.WriteByte(crossword[r+i*dir.y][c+i*dir.x])
	}
	return sb.String()
}

func main() {
	crossword := getInput("input.txt")
	total := processCrossword(crossword)
	fmt.Println(total)

}
