package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := getInput("input.txt")
	// data := getInput("smallInput.txt")
	processedData := processData(data)
	// fmt.Println(processedData)
	stones := multiplyStones(processedData, 25)
	fmt.Println(len(stones))
	// fmt.Println(stones)
}

func getInput(file string) []string {
	contents, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	contentsStr := string(contents)

	return strings.Fields(contentsStr)
}

func processData(data []string) []int {
	processedData := []int{}
	for _, datum := range data {
		num, err := strconv.Atoi(datum)
		if err != nil {
			log.Fatal(err)
		}
		processedData = append(processedData, num)
	}

	return processedData
}

func multiplyStones(data []int, blinks int) []int {
	stones := data
	for i := 0; i < blinks; i++ {
		// fmt.Printf("Blink: %d\n%v\n", i, stones)
		fmt.Println(i)
		stones = blink(stones)
	}

	return stones
}

func blink(stones []int) []int {
	//rules:
	//1. If 0, becomes 1
	//2. If an even num of digits split the digits remove leading 0s
	//3. If neither of the above apply, multiply by 2024
	newStones := []int{}
	for _, stone := range stones {
		stoneString := strconv.Itoa(stone)
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if len(stoneString)%2 == 0 {
			left := stoneString[:len(stoneString)/2]
			right := stoneString[len(stoneString)/2:]
			// fmt.Printf("Left: %v, Right: %v\n", left, right)
			l, err := strconv.Atoi(left)
			if err != nil {
				log.Fatal(err)
			}
			newStones = append(newStones, l)
			r, err := strconv.Atoi(right)
			if err != nil {
				log.Fatal(err)
			}
			newStones = append(newStones, r)
			// fmt.Println(stones)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}
	return newStones
}
