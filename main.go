package main

import (
	"flag"
	"fmt"

	"github.com/junijland/aoc2024/day01"
	"github.com/junijland/aoc2024/day02"
	"github.com/junijland/aoc2024/day03"
	"github.com/junijland/aoc2024/day04"
)

func main() {
	day := flag.Int("day", 0, "day of advent of code to solve")
	flag.Parse()

	if *day == 0 {
		fmt.Println("Please provide a day number using -day flag")
		flag.Usage()
		return
	}

	switch *day {
	case 1:
		day01.Solve()
	case 2:
		day02.Solve()
	case 3:
		day03.Solve()
	case 4:
		day04.Solve()
	default:
		fmt.Printf("Day %d not implemented yet\n", *day)
	}
}
