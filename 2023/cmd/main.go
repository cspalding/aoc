package main

import (
	"fmt"
	"os"

	"github.com/cspalding/aoc/2023/internal/day1"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong number of args.\nUsage:\n\tgo run main.go DAY FILENAME")
		return
	}

	day := os.Args[1]
	filename := os.Args[2]

	switch day {
	case "1":
		fmt.Printf("Day 1 Part 1 Solution: %d\n", day1.Day1Part1(filename))
		fmt.Printf("Day 1 Part 2 Solution: %d\n", day1.Day1Part2(filename))
	default:
		fmt.Printf("Solution not implemented for day %s", day)
	}
}
