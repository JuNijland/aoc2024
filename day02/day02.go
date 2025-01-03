package day02

import (
	"fmt"
	"math"
	"strings"

	"github.com/junijland/aoc2024/utils"
)

func Solve() {
	input := utils.ReadInputFile("day02/input.txt")
	reports := parseNumbers(input)

	part1 := solvePart1(reports)
	part2 := solvePart2(reports)
	fmt.Printf("Day 02 - Part 1: %v\n", part1)
	fmt.Printf("Day 02 - Part 2: %v\n", part2)

}

func parseNumbers(input []string) [][]int {
	var reports [][]int

	for _, line := range input {
		var levels []int
		var num int

		// Read all integers from the line
		for _, field := range strings.Fields(line) {
			fmt.Sscanf(field, "%d", &num)
			levels = append(levels, num)
		}

		reports = append(reports, levels)
	}

	return reports
}

func solvePart1(reports [][]int) int {
	var safeReports int = 0
	for _, report := range reports {
		if isReportSafePart1(report) {
			safeReports += 1
		}
	}
	return safeReports
}

func solvePart2(reports [][]int) int {
	var safeReports int = 0
	for _, report := range reports {
		if isReportSafePart2(report) {
			safeReports += 1
		}
	}
	return safeReports
}

func isReportSafePart1(report []int) bool {
	if len(report) < 2 {
		return true
	}

	asc := report[0] < report[1]

	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i] - report[i-1]))
		if diff < 1 || diff > 3 || (report[i-1] < report[i]) != asc {
			return false
		}
	}

	return true
}

func isReportSafePart2(report []int) bool {
	if len(report) < 2 {
		return true
	}

	// First check if it's safe without removing anything
	if isSequenceSafe(report) {
		return true
	}

	// Try removing each number one at a time
	for i := 0; i < len(report); i++ {
		tempReport := make([]int, 0, len(report)-1)
		tempReport = append(tempReport, report[:i]...)
		tempReport = append(tempReport, report[i+1:]...)

		if isSequenceSafe(tempReport) {
			return true
		}
	}

	return false
}

func isSequenceSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	// Check if sequence is monotonic (all ascending or all descending)
	ascending := true
	descending := true

	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i] - report[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}

		if report[i] <= report[i-1] {
			ascending = false
		}
		if report[i] >= report[i-1] {
			descending = false
		}
	}

	return ascending || descending
}
