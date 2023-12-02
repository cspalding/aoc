package day1

import (
	"bufio"
	"fmt"
	"os"
)

func GetFirstAndLastDigit(line string) int {
	firstDigit := 0
	lastDigit := 0
	for _, char := range line {
		res := int(char - '0')

		if res < 10 && firstDigit == 0 {
			firstDigit = res
		} else if res < 10 {
			lastDigit = res
		}
	}

	if lastDigit == 0 {
		lastDigit = firstDigit
	}

	return (10 * firstDigit) + lastDigit
}

func Day1Part1(filepath string) int {
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		sum += GetFirstAndLastDigit(scanner.Text())
	}
	return sum
}
