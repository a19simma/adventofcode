package adventofcode

import (
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestSimleSeedLocation(t *testing.T) {
	seeds, maps := GetMaps(input)
	result := GetLowestLocation(seeds, maps)
	assert.Equal(t, 35, result)
}

func TestSeedLocation(t *testing.T) {
	puzzleData, _ := os.ReadFile("5.txt")
	seeds, maps := GetMaps(string(puzzleData))
	result := GetLowestLocation(seeds, maps)
	assert.Equal(t, 31599214, result)
}
func TestSimpleSeedRangeLocation(t *testing.T) {
	seeds, maps := GetRangeSeedMaps(input)
	result := GetLowestLocation(seeds, maps)
	assert.Equal(t, 46, result)
}

func TestSeedRangeLocation(t *testing.T) {
	puzzleData, _ := os.ReadFile("5.txt")
	seeds, maps := GetRangeSeedMaps(string(puzzleData))
	result := GetLowestLocation(seeds, maps)
	assert.Equal(t, 20358600, result)
}

type Mapping struct {
	destination, source, rangeLength int
}

type Pair struct {
	start, length int
}

func GetRangeSeedMaps(data string) ([]int, [][]Mapping) {
	maps := [][]Mapping{}

	sections := strings.Split(data, "\n\n")
	seedRanges := strings.Split(sections[0], " ")[1:]

	length := 0

	pairs := []Pair{}
	for i := 0; i < len(seedRanges)-1; i += 2 {
		start, _ := strconv.Atoi(seedRanges[i])
		l, _ := strconv.Atoi(seedRanges[i+1])
		pairs = append(pairs, Pair{start: start, length: l})
		length += l
	}

	seeds := make([]int, length)
	j := 0
	for _, p := range pairs {
		for k := p.start; k < p.start+p.length; k++ {
			seeds[j] = k
			j++
		}
	}
	for _, s := range sections[1:] {
		ranges := strings.Split(s, "\n")[1:]
		tmpMapArray := []Mapping{}
		for _, r := range ranges {
			values := strings.Split(r, " ")
			destination, _ := strconv.Atoi(values[0])
			source, _ := strconv.Atoi(values[1])
			rangeLength, _ := strconv.Atoi(values[2])
			tmpMapArray = append(tmpMapArray, Mapping{destination: destination, source: source, rangeLength: rangeLength})
		}
		maps = append(maps, tmpMapArray)
	}

	return seeds, maps
}

func GetLowestLocation(seeds []int, maps [][]Mapping) int {
	for i, s := range seeds {
		for _, m := range maps {
			for _, mapping := range m {
				if mapping.source > s {
					continue
				}
				if mapping.source+mapping.rangeLength < s {
					continue
				}
				t := mapping.destination + s - mapping.source
				seeds[i] = t
				s = t
				break
			}

		}
	}
	return slices.Min(seeds)
}

func GetMaps(data string) ([]int, [][]Mapping) {
	maps := [][]Mapping{}
	seeds := []int{}

	sections := strings.Split(data, "\n\n")
	for _, s := range strings.Fields(sections[0])[1:] {
		i, _ := strconv.Atoi(s)
		seeds = append(seeds, i)
	}
	for _, s := range sections[1:] {
		ranges := strings.Split(s, "\n")[1:]
		tmpMapArray := []Mapping{}
		for _, r := range ranges {
			values := strings.Split(r, " ")
			destination, _ := strconv.Atoi(values[0])
			source, _ := strconv.Atoi(values[1])
			rangeLength, _ := strconv.Atoi(values[2])
			tmpMapArray = append(tmpMapArray, Mapping{destination: destination, source: source, rangeLength: rangeLength})
		}
		maps = append(maps, tmpMapArray)
	}

	return seeds, maps
}
