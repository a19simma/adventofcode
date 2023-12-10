package adventofcode

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleReport(t *testing.T) {
	input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	report := ParseReport(input)
	result := ExtrapolateSum(report)
	assert.Equal(t, 114, result)
}
func TestReport(t *testing.T) {
	input, _ := os.ReadFile("9.txt")
	report := ParseReport(string(input))
	result := ExtrapolateSum(report)
	assert.Equal(t, 1762065988, result)
}

func TestSimpleReverseReport(t *testing.T) {
	input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	report := ParseReport(input)
	result := ExtrapolateSum2(report)
	assert.Equal(t, 2, result)
}
func TestReverseReport(t *testing.T) {
	input, _ := os.ReadFile("9.txt")
	report := ParseReport(string(input))
	result := ExtrapolateSum2(report)
	assert.Equal(t, 1066, result)
}

func ExtrapolateSum2(report [][]int) int {
	result := 0
	for _, l := range report {
		if len(l) == 0 {
			continue
		}
		tmp := [][]int{}
		tmp = append(tmp, l)
		stop := false
		cur := l
		fmt.Println(cur)
		for !stop {
			cur = nextDif(cur)
			tmp = append(tmp, cur)
			fmt.Println(cur)
			stop = allZero(cur)
		}
		next := 0
		for i := len(tmp) - 1; i > 0; i-- {
			b := tmp[i-1][0]
			next = b - next
		}
		fmt.Println(next)
		result += next

	}
	return result
}

func ExtrapolateSum(report [][]int) int {
	result := 0
	for _, l := range report {
		if len(l) == 0 {
			continue
		}
		tmp := [][]int{}
		tmp = append(tmp, l)
		stop := false
		cur := l
		fmt.Println(cur)
		for !stop {
			cur = nextDif(cur)
			tmp = append(tmp, cur)
			fmt.Println(cur)
			stop = allZero(cur)
		}
		next := 0
		for i := len(tmp) - 1; i > 0; i-- {
			b := tmp[i-1][len(tmp[i-1])-1]
			next += b
		}
		fmt.Println(next)
		result += next

	}
	return result
}

func allZero(a []int) bool {
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}

func nextDif(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	b := make([]int, len(a)-1)
	for i := 0; i < len(a)-1; i++ {
		b[i] = a[i+1] - a[i]
	}
	return b
}

func ParseReport(input string) [][]int {
	result := [][]int{}
	fmt.Println("Report:")
	for _, s := range strings.Split(input, "\n") {
		line := []int{}
		for _, f := range strings.Fields(s) {
			i, _ := strconv.Atoi(f)
			line = append(line, i)
		}
		fmt.Println(line)
		result = append(result, line)
	}
	fmt.Println("--- END ---")
	return result
}
