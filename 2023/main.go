package main

import (
	"fmt"
	"os"
)

func main() {
	day := os.Args[1]

	switch day {
	case "1":
		Day1Part1()
	default:
		fmt.Printf("Solution not implemented for day %s", day)
	}
}
