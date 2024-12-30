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
	stones := multiplyStones(processedData, 75)
	fmt.Println(stones)
}

func getInput(file string) []string {
	contents, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	contentsStr := string(contents)

	return strings.Fields(contentsStr)
}

func processData(data []string) map[int]int {
	processedData := map[int]int{}
	for _, datum := range data {
		num, err := strconv.Atoi(datum)
		if err != nil {
			log.Fatal(err)
		}
		processedData[num]++
	}

	return processedData
}

func multiplyStones(data map[int]int, blinks int) int {
	stones := data
	for i := 0; i < blinks; i++ {
		// fmt.Printf("Blink: %d\n%v\n", i, stones)
		stones = blink(stones)
	}

	total := 0
	for _, val := range stones {
		total += val
	}

	return total
}

func blink(stones map[int]int) map[int]int {
	//rules:
	//1. If 0, becomes 1
	//2. If an even num of digits split the digits remove leading 0s
	//3. If neither of the above apply, multiply by 2024
	newStones := map[int]int{}
	for key, val := range stones {
		stoneString := strconv.Itoa(key)
		if key == 0 {
			newStones[1] += val
		} else if len(stoneString)%2 == 0 {
			left := stoneString[:len(stoneString)/2]
			right := stoneString[len(stoneString)/2:]
			// fmt.Printf("Left: %v, Right: %v\n", left, right)
			l, err := strconv.Atoi(left)
			if err != nil {
				log.Fatal(err)
			}
			newStones[l] += val
			r, err := strconv.Atoi(right)
			if err != nil {
				log.Fatal(err)
			}
			newStones[r] += val
		} else {
			newKey := key * 2024
			newStones[newKey] += val
		}
	}

	return newStones
}
