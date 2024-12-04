package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput(file string) string {
	contents, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}

func processData(data string) int {
	total := 0
	splitContents := strings.Split(data, "mul(")

	for i := 0; i < len(splitContents); i++ {
		total += processStrings(splitContents[i])
	}
	return total
}

func processStrings(data string) int {
	before, _, found := strings.Cut(data, ")")
	if !found {
		return 0
	}

	splitData := strings.Split(before, ",")

	if len(splitData) != 2 {
		return 0
	}
	x, err := strconv.Atoi(splitData[0])
	if err != nil {
		return 0
	}
	if x > 999 {
		return 0
	}
	y, err := strconv.Atoi(splitData[1])
	if err != nil {
		return 0
	}
	if y > 999 {
		return 0
	}
	return x * y

}

func main() {
	contents := getInput("input.txt")
	total := processData(contents)

	fmt.Println(total)
}
