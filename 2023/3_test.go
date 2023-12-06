package adventofcode

import (
	"os"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestSimpleEngineSchematic(t *testing.T) {
	data := `467..114..
...*......
..35*.633.
......#...
617.......
*....+.58.
..592.....
......*755
..........
.664%598..`
	result := GetAdjacentSymbolNumbers(data)
	assert.Equal(t, 4361, result)
}

func TestEngineSchematic(t *testing.T) {
	data, _ := os.ReadFile("3.txt")
	result := GetAdjacentSymbolNumbers(string(data))
	assert.Equal(t, 531561, result)
}
func TestSimpleGearRatio(t *testing.T) {
	data := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	result := GetGearRatio(data)
	assert.Equal(t, 467835, result)
}
func TestGearRatio(t *testing.T) {
	data, _ := os.ReadFile("3.txt")
	result := GetGearRatio(string(data))
	assert.Equal(t, 83279367, result)
}

func GetGearRatio(data string) int {
	result := 0
	possibleGears := []Coordinate{}
	lines := strings.Split(data, "\n")
	maxes := Coordinate{x: len(lines[0]) - 1, y: len(lines) - 1}
	for y, l := range lines {
		for x, r := range l {
			if r == '*' {
				possibleGears = append(possibleGears, Coordinate{x: x, y: y})
			}
		}
	}
	parts := GetParts(lines)
	for _, g := range possibleGears {
		matches := 0
		first := 0
		second := 0
		for _, p := range parts {
			minStart := p.start - 1
			maxEnd := p.end + 1
			minY := p.y - 1
			maxY := p.y + 1
			if minStart < 0 {
				minStart = 0
			}
			if maxEnd > maxes.x {
				maxEnd = maxes.x
			}
			if minY < 0 {
				minY = 0
			}
			if maxY > maxes.y {
				maxY = maxes.y
			}
			if (minStart <= g.x && g.x <= maxEnd) && (minY <= g.y && g.y <= maxY) {
				matches += 1
				if first == 0 {
					first = p.value
				} else if second == 0 {
					second = p.value
				}
			}
		}
		if matches == 2 {
			result += first * second
		}
		matches = 0
		first = 0
		second = 0
	}
	return result
}

func GetParts(lines []string) []Part {
	parts := []Part{}
	for y, l := range lines {
		number := ""
		var start, end int
		first := false
		for x, r := range l {
			if unicode.IsDigit(r) {
				if !first {
					start = x
					first = true
				}
				number += string(r)
				if x == len(l)-1 {
					n, _ := strconv.Atoi(number)
					parts = append(parts, Part{start: start, end: end, y: y, value: n})
					number = ""
					first = false
				}
			} else {
				if number == "" {
					continue
				}
				end = x - 1
				n, _ := strconv.Atoi(number)
				parts = append(parts, Part{start: start, end: end, y: y, value: n})
				number = ""
				first = false
			}
		}
	}
	return parts
}

type Part struct {
	start, end, y, value int
}

type Coordinate struct {
	x, y int
}

func GetAdjacentSymbolNumbers(data string) int {
	result := 0
	symbolCoordinates := []Coordinate{}
	lines := strings.Split(data, "\n")
	maxes := Coordinate{x: len(lines[0]) - 1, y: len(lines) - 1}
	for y, l := range lines {
		for x, r := range l {
			if !unicode.IsDigit(r) && r != '.' {
				symbolCoordinates = append(symbolCoordinates, Coordinate{x: x, y: y})
			}
		}
	}
	for y, l := range lines {
		number := ""
		var start, end int
		first := false
		for x, r := range l {
			if unicode.IsDigit(r) {
				if !first {
					start = x
					first = true
				}
				number += string(r)
				if x == len(l)-1 {
					if CoordinatesAdjacent(start, x, y, symbolCoordinates, maxes) {
						n, _ := strconv.Atoi(number)
						result += n
					}
					number = ""
					first = false
				}
			} else {
				if number == "" {
					continue
				}
				end = x - 1
				if CoordinatesAdjacent(start, end, y, symbolCoordinates, maxes) {
					n, _ := strconv.Atoi(number)
					result += n
				}
				number = ""
				first = false
			}
		}
	}
	return result
}

func CoordinatesAdjacent(start, end, y int, symbolCoordinates []Coordinate, maxes Coordinate) bool {
	minStart := start - 1
	maxEnd := end + 1
	minY := y - 1
	maxY := y + 1
	if minStart < 0 {
		minStart = 0
	}
	if maxEnd > maxes.x {
		maxEnd = maxes.x
	}
	if minY < 0 {
		minY = 0
	}
	if maxY > maxes.y {
		maxY = maxes.y
	}
	for _, c := range symbolCoordinates {
		if (minStart <= c.x && c.x <= maxEnd) && (minY <= c.y && c.y <= maxY) {
			return true
		}
	}
	return false
}
