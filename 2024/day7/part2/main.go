package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := getInput("input.txt")

	equations := processData(data)

	validEquations := getValidEquations(equations)

	sum := sumEquations(validEquations)

	fmt.Println(sum)
	//last answer 145398187529596

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

func processData(data []string) map[int][]int {
	equations := map[int][]int{}
	for _, datum := range data {
		b, a, f := strings.Cut(datum, ":")
		if !f {
			fmt.Print("Invalid Line: ")
			fmt.Println(datum)
		}

		key, err := strconv.Atoi(b)
		if err != nil {
			log.Fatal(err)
		}
		fields := strings.Fields(a)
		nums := []int{}
		for _, val := range fields {
			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Print("Invalid Line: ")
				fmt.Println(val)
				log.Fatal(err)
			}
			nums = append(nums, num)
		}
		equations[key] = nums

	}

	return equations
}

func getValidEquations(equations map[int][]int) []int {
	validTotals := []int{}

	for total, nums := range equations {
		if isValidEquation(total, nums) {
			validTotals = append(validTotals, total)
		}
	}
	return validTotals
}

func isValidEquation(target int, nums []int) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}

	l := len(nums) - 1
	if isValidEquation(target-nums[l], nums[:l]) {
		return true
	}
	if target%nums[l] == 0 {
		if isValidEquation(target/nums[l], nums[:l]) {
			return true
		}
	}
	//concat last two numbers
	targetStr := strconv.Itoa(target)
	lastStr := strconv.Itoa(nums[l])
	if !strings.HasSuffix(targetStr, lastStr) {
		return false
	}
	targetStr = strings.TrimSuffix(targetStr, lastStr)
	newTarget, err := strconv.Atoi(targetStr)
	if err != nil {
		return false
	}
	if isValidEquation(newTarget, nums[:l]) {
		return true
	}
	return false
}

func sumEquations(validTotals []int) int {
	total := 0
	for _, val := range validTotals {
		total += val
	}
	return total
}
