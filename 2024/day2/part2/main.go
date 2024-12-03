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

func getInput(file string) [][]int {
	reports := [][]int{}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		report := []int{}
		reportString := strings.Fields(scanner.Text())
		for i := 0; i < len(reportString); i++ {
			val, err := strconv.Atoi(reportString[i])
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, val)

		}
		reports = append(reports, report)
	}
	return reports
}

func processReports(reports [][]int) int {
	safeReports := 0
	for i := 0; i < len(reports); i++ {
		if isReportSafe(reports[i]) {
			safeReports++
		} else if isReportSafeTolerant(reports[i]) {
			safeReports++
		}
	}

	return safeReports
}

func isReportSafe(report []int) bool {
	if report[0] > report[1] {
		slices.Reverse(report)
	}
	if !slices.IsSorted(report) {
		return false
	}
	for i := 1; i < len(report); i++ {
		if !isDiffSafe(report[i], report[i-1]) {
			return false
		}
	}

	return true

}

func isReportSafeTolerant(report []int) bool {
	if isReportSafe(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		newReport := removeIndex(report, i)
		if isReportSafe(newReport) {
			return true
		}
	}

	return false

}

func removeIndex(report []int, index int) []int {
	newReport := make([]int, 0, len(report)-1)
	for i := 0; i < len(report); i++ {
		if i != index {
			newReport = append(newReport, report[i])
		}
	}
	return newReport
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func isDiffSafe(x int, y int) bool {
	val := abs(x - y)
	return val <= 3 && val > 0
}

func main() {
	reports := getInput("input.txt")
	numSafeReports := processReports(reports)
	fmt.Println(numSafeReports)

}
