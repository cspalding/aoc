package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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

func Day1Part2(filepath string) int {
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += findFirst(line)
		sum += findLast(line)
	}
	return sum
}

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

func findFirst(input string) int {
	pattern := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	match := pattern.Find([]byte(input))

	if len(match) < 1 {
		fmt.Printf("input '%s' had no matches!", input)
		return -1
	}

	first := stringToNumber(string(match))

	if first == -1 {
		fmt.Printf("ERROR FINDING FIRST DIGIT FROM INPUT '%s'\n", input)
		return -1
	}
	return (first * 10)
}

func findLast(input string) int {
	reversed := reverse(input)
	pattern := regexp.MustCompile(`(\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)
	match := pattern.Find([]byte(reversed))

	// fmt.Printf("FOUND MATCH '%s' FROM INPUT '%s'\n", match, input)
	return stringToNumber(
		reverse(string(match)),
	)
}

func reverse(input string) string {
	res := ""
	for idx := len(input) - 1; idx > -1; idx-- {
		res += string(input[idx])
	}
	return res
}

func stringToNumber(input string) int {
	if res, err := strconv.Atoi(input); err != nil {
		switch input {
		case "one":
			return 1
		case "two":
			return 2
		case "three":
			return 3
		case "four":
			return 4
		case "five":
			return 5
		case "six":
			return 6
		case "seven":
			return 7
		case "eight":
			return 8
		case "nine":
			return 9
		default:
			return -1
		}
	} else {
		return res
	}
}
