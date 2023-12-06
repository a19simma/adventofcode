package adventofcode

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleBoatRace(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`
	races := GetRaces(input)
	result := FindWinning(races)
	assert.Equal(t, 288, result)
}

func TestBoatRace(t *testing.T) {
	input := `Time:        52     94     75     94
Distance:   426   1374   1279   1216`
	races := GetRaces(input)
	result := FindWinning(races)
	assert.Equal(t, 0, result)
}
func TestBigBoatRace(t *testing.T) {
	input := `Time:        52947594
Distance:   426137412791216`
	races := GetRaces(input)
	result := FindWinning(races)
	assert.Equal(t, 0, result)
}

func GetRaces(input string) []Race {
	races := []Race{}
	lines := strings.Split(input, "\n")
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]
	for i, t := range times {
		ti, _ := strconv.Atoi(t)
		di, _ := strconv.Atoi(distances[i])
		races = append(races, Race{time: ti, distance: di})
	}
	return races
}

func FindWinning(races []Race) int {
	result := 0
	for _, r := range races {
		var highest, lowest int
		for i := 0; i < r.time; i++ {
			distance := i * (r.time - i)
			if distance > r.distance {
				lowest = i
				break
			}
		}

		for i := r.time; i > lowest; i-- {
			distance := i * (r.time - i)
			if distance > r.distance {
				highest = i + 1
				break
			}
		}
		if result == 0 {
			result = highest - lowest
		} else {
			result *= highest - lowest
		}
	}
	return result
}

type Race struct {
	time, distance int
}
