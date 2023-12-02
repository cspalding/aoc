package day1

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
	return (10 * firstDigit) + lastDigit
}

func Day1Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += GetFirstAndLastDigit(line)
	}
	return sum
}
