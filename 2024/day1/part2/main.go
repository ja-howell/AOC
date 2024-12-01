package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput2(file string) (nums []int, freq map[int]int) {
	contents, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	contentsStr := string(contents)

	splitContents := strings.Split(contentsStr, "\n")

	freq = map[int]int{}
	for _, line := range splitContents {
		fields := strings.Fields(line)
		leftField, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		rightField, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal()
		}
		nums = append(nums, leftField)
		freq[rightField]++

	}

	return nums, freq
}

// func getInput(file string) (nums []int, freq map[int]int) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)

// 	freq = map[int]int{}

// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 		fields := strings.Fields(scanner.Text())
// 		leftField, err := strconv.Atoi(fields[0])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		rightField, err := strconv.Atoi(fields[1])
// 		if err != nil {
// 			log.Fatal()
// 		}
// 		nums = append(nums, leftField)
// 		freq[rightField]++

// 	}

// 	return nums, freq
// }

func calcSimilarityScore(left []int, right map[int]int) int {
	score := 0
	for i := 0; i < len(left); i++ {
		score += left[i] * right[left[i]]
	}
	return score
}

func main() {
	nums, freq := getInput2("input.txt")
	score := calcSimilarityScore(nums, freq)
	fmt.Println(score)
}
