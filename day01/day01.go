package day01

import (
	"fmt"
	"math"
	"sort"

	"github.com/junijland/aoc2024/utils"
)

func Solve() {
	input := utils.ReadInputFile("day01/input.txt")
	leftNums, rightNums := parseNumbers(input)

	part1 := solvePart1(leftNums, rightNums)
	part2 := solvePart2(leftNums, rightNums)

	fmt.Printf("Day 01 - Part 1: %v\n", part1)
	fmt.Printf("Day 01 - Part 2: %v\n", part2)
}

func parseNumbers(input []string) ([]int, []int) {
	numLines := len(input)

	leftNums := make([]int, numLines)
	rightNums := make([]int, numLines)

	for i, line := range input {
		var left, right int
		fmt.Sscanf(line, "%d %d", &left, &right)
		leftNums[i] = left
		rightNums[i] = right
	}

	return leftNums, rightNums
}

func solvePart1(leftNums []int, rightNums []int) int {
	// Sort both lists in ascending order
	sort.Ints(leftNums)
	sort.Ints(rightNums)

	diffTotal := 0
	// Get diff between each pair in the list and sum it
	for i := 0; i < len(leftNums); i++ {
		diff := math.Abs(float64(leftNums[i] - rightNums[i]))
		diffTotal += int(diff)
	}

	return diffTotal
}

func solvePart2(leftNums []int, rightNums []int) int {
	// Create a map to store frequencies of numbers in rightNums
	frequencies := make(map[int]int)

	// Count occurrences of each number in rightNums
	for _, num := range rightNums {
		frequencies[num]++
	}

	totalSimilarityScore := 0
	for _, num := range leftNums {
		totalSimilarityScore += num * frequencies[num]
	}

	return totalSimilarityScore
}
