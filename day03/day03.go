package day03

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/junijland/aoc2024/utils"
)

func Solve() {
	input := utils.ReadInputFile("day03/input.txt")

	part1 := solvePart1(input)
	part2 := solvePart2(input)
	fmt.Printf("Day 03 - Part 1: %v\n", part1)
	fmt.Printf("Day 03 - Part 2: %v\n", part2)

}

func solvePart1(input []string) int {
	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	sum := 0

	for _, line := range input {
		matches := mulRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}

	return sum
}

func solvePart2(input []string) int {
	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	sum := 0
	enabled := true

	for _, line := range input {
		pos := 0
		for pos < len(line) {
			doMatch := doRegex.FindStringIndex(line[pos:])
			dontMatch := dontRegex.FindStringIndex(line[pos:])
			mulMatch := mulRegex.FindStringSubmatchIndex(line[pos:])

			nextPos := len(line)
			if doMatch != nil && (doMatch[0] < nextPos) {
				nextPos = doMatch[0]
			}
			if dontMatch != nil && (dontMatch[0] < nextPos) {
				nextPos = dontMatch[0]
			}
			if mulMatch != nil && (mulMatch[0] < nextPos) {
				nextPos = mulMatch[0]
			}

			if nextPos == len(line) {
				break
			}

			remainder := line[pos:]
			if doMatch != nil && doMatch[0] == nextPos {
				enabled = true
				pos += doMatch[1]
			} else if dontMatch != nil && dontMatch[0] == nextPos {
				enabled = false
				pos += dontMatch[1]
			} else if mulMatch != nil && mulMatch[0] == nextPos {
				if enabled {
					matches := mulRegex.FindStringSubmatch(remainder)
					num1, _ := strconv.Atoi(matches[1])
					num2, _ := strconv.Atoi(matches[2])
					sum += num1 * num2
				}
				pos += mulMatch[1]
			}
		}
	}

	return sum
}
