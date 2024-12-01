package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getInput(file string) (left []int, right []int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fields := strings.Fields(scanner.Text())
		leftField, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		rightField, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal()
		}
		left = append(left, leftField)
		right = append(right, rightField)
	}

	return left, right
}

func calcTotalDistance(left []int, right []int) int {
	totDist := 0
	for i := 0; i < len(left); i++ {
		totDist += abs(left[i] - right[i])
	}

	return totDist
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	left, right := getInput("input.txt")
	slices.Sort(left)
	slices.Sort(right)
	totDist := calcTotalDistance(left, right)
	fmt.Println(totDist)
}
