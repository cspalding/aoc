package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		Name     string
		Input    string
		Expected int
	}{
		{Name: "Simplest Input", Input: "1abc2", Expected: 12},
		{Name: "Digits in Middle of String", Input: "pqr3stu8vwx", Expected: 38},
		{Name: "More than two digits", Input: "a1b2c3d4e5f", Expected: 15},
		{Name: "Only one digit", Input: "treb7uchet", Expected: 77},
	}

	for _, tc := range cases {
		actual := GetFirstAndLastDigit(tc.Input)
		assert.Equal(t, tc.Expected, actual, tc.Name)
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		Name     string
		Input    string
		Expected string
	}{
		{Name: "Reverse Test 1", Input: "two1nine", Expected: "enin1owt"},
		{Name: "Reverse Test 2", Input: "5threeeightwor", Expected: "rowthgieeerht5"},
	}

	for _, tc := range cases {
		actual := reverse(tc.Input)
		assert.Equal(t, tc.Expected, actual, tc.Name)
	}
}

func TestFindLast(t *testing.T) {
	cases := []struct {
		Name     string
		Input    string
		Expected int
	}{
		{Name: "Find Last Test 1", Input: "two1nine", Expected: 9},
		{Name: "Find Last Test 2", Input: "5threeeightwor", Expected: 2},
	}

	for _, tc := range cases {
		actual := findLast(tc.Input)
		assert.Equal(t, tc.Expected, actual, tc.Name)
	}
}
