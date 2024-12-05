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

func main() {
	data := getInput("input.txt")
	rules, updates := processData(data)
	// fmt.Println(rules)
	// for _, v := range updates {
	// 	fmt.Println(v)
	// }
	total := calculateTotal(rules, updates)

	fmt.Println(total)
}

func getInput(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	data := []string{}

	for s.Scan() {
		data = append(data, s.Text())
	}

	return data

}

func processData(data []string) (rules map[int][]int, updates [][]int) {
	rules = map[int][]int{}
	for _, v := range data {
		if strings.Contains(v, "|") {
			splitVal := strings.Split(v, "|")
			key := Atoi(splitVal[0])
			rule := Atoi(splitVal[1])
			rules[key] = append(rules[key], rule)

		} else if strings.Contains(v, ",") {
			splitVal := strings.Split(v, ",")
			temp := []int{}
			for _, num := range splitVal {
				temp = append(temp, Atoi(num))
			}

			updates = append(updates, temp)
		}
	}
	return rules, updates
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func calculateTotal(rules map[int][]int, updates [][]int) int {
	total := 0
	for i := 0; i < len(updates); i++ {
		if updateIsValid(rules, updates[i]) {
			total += updates[i][len(updates[i])/2]
		}
	}
	return total
}

func updateIsValid(rules map[int][]int, update []int) bool {
	for i, v := range update {
		for l := i - 1; l >= 0; l-- {
			if slices.Contains(rules[v], update[l]) {
				return false
			}
		}
	}
	return true
}
