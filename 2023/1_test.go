package adventofcode

import (
	"os"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestCalibration(t *testing.T) {
	data, _ := os.ReadFile("1.txt")
	result := Calibration(strings.Split(string(data), "\n"))
	assert.Equal(t, 52834, result, "should be equal") // 53334 for the first task
}

func Calibration(values []string) int {
	result := 0
	numbers := map[string]rune{"one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9'}
	for _, v := range values {
		var firstDigit, lastDigit string

		for i, r := range []rune(v) {
			var threeLetter, fourLetter, fiveLetter string
			if i+2 < len(v) {
				threeLetter = v[i : i+3]
			}
			if i+3 < len(v) {
				fourLetter = v[i : i+4]
			}
			if i+4 < len(v) {
				fiveLetter = v[i : i+5]
			}

			val, ok := numbers[threeLetter]
			if ok {
				r = val
			}
			val, ok = numbers[fourLetter]
			if ok {
				r = val
			}
			val, ok = numbers[fiveLetter]
			if ok {
				r = val
			}

			if !unicode.IsDigit(r) {
				continue
			}

			s := string(r)
			if firstDigit == "" {
				firstDigit = s
			}
			lastDigit = s
		}
		t, _ := strconv.Atoi(firstDigit + lastDigit)
		result += t
	}
	return result
}
