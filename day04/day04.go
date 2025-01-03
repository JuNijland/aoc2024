package day04

import (
	"fmt"

	"github.com/junijland/aoc2024/utils"
)

// Direction represents possible movement directions in the grid
type Direction struct {
	dx, dy int
}

// All possible 8-directional movements
var directions = []Direction{
	{-1, -1}, {-1, 0}, {-1, 1}, // top-left, top, top-right
	{0, -1}, {0, 1}, // left, right
	{1, -1}, {1, 0}, {1, 1}, // bottom-left, bottom, bottom-right
}

func Solve() {
	input := utils.ReadInputFile("day04/input.txt")

	part1 := solvePart1(input)
	part2 := solvePart2(input)
	fmt.Printf("Day 04 - Part 1: %v\n", part1)
	fmt.Printf("Day 04 - Part 2: %v\n", part2)
}

func solvePart1(input []string) int {
	grid := createGrid(input)
	return countXMASOccurrences(grid)
}

func solvePart2(input []string) int {
	grid := createGrid(input)
	return countXMASPatternsFromA(grid)
}

func createGrid(input []string) [][]rune {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}
	return grid
}

func countXMASOccurrences(grid [][]rune) int {
	count := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'X' {
				count += findXMASFromPosition(grid, x, y)
			}
		}
	}
	return count
}

func findXMASFromPosition(grid [][]rune, x, y int) int {
	count := 0
	for _, dir := range directions {
		if canFormXMAS(grid, x, y, dir) {
			count++
		}
	}
	return count
}

func canFormXMAS(grid [][]rune, x, y int, dir Direction) bool {
	target := "XMAS"
	rows, cols := len(grid), len(grid[0])

	for i, char := range target {
		newX, newY := x+(dir.dx*i), y+(dir.dy*i)

		if !isValidPosition(newX, newY, rows, cols) {
			return false
		}
		if grid[newY][newX] != char {
			return false
		}
	}
	return true
}

func isValidPosition(x, y, rows, cols int) bool {
	return x >= 0 && x < cols && y >= 0 && y < rows
}

func countXMASPatternsFromA(grid [][]rune) int {
	count := 0
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if grid[y][x] == 'A' && findXPatternFromA(grid, x, y) {
				count++
			}
		}
	}
	return count
}

func findXPatternFromA(grid [][]rune, x, y int) bool {
	// Check all 4 combinations
	return (
	// 1. Top-left to bottom-right: MAS, top-right to bottom-left: MAS
	(grid[y-1][x-1] == 'M' && grid[y+1][x+1] == 'S' &&
		grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S') ||

		// 2. Top-left to bottom-right: MAS, top-right to bottom-left: SAM
		(grid[y-1][x-1] == 'M' && grid[y+1][x+1] == 'S' &&
			grid[y-1][x+1] == 'S' && grid[y+1][x-1] == 'M') ||

		// 3. Top-left to bottom-right: SAM, top-right to bottom-left: MAS
		(grid[y-1][x-1] == 'S' && grid[y+1][x+1] == 'M' &&
			grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S') ||

		// 4. Top-left to bottom-right: SAM, top-right to bottom-left: SAM
		(grid[y-1][x-1] == 'S' && grid[y+1][x+1] == 'M' &&
			grid[y-1][x+1] == 'S' && grid[y+1][x-1] == 'M'))
}
