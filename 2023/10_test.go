package adventofcode

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimlePipeMap(t *testing.T) {
	input := `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`
	schematic, sx, sy := GetSchematic(input)
	steps := LengthOfPipes(schematic, sx, sy)
	assert.Equal(t, 8, steps)
}

func TestPipeMap(t *testing.T) {
	input, _ := os.ReadFile("10.txt")
	schematic, sx, sy := GetSchematic(string(input))
	steps := LengthOfPipes(schematic, sx, sy)
	assert.Equal(t, 6701, steps)
}

func GetSchematic(input string) (schematic [][]Pipe, sx, sy int) {
	for i, l := range strings.Split(input, "\n") {
		line := make([]Pipe, len(l))
		for j, p := range l {
			if p == '.' {
				fmt.Print(".")
				continue
			}
			if p == 'S' {
				sx = j
				sy = i
			}
			line[j] = pipemap[p]
			fmt.Print(string(p))
		}
		fmt.Print("\n")
		schematic = append(schematic, line)
	}
	return schematic, sx, sy
}

func LengthOfPipes(schematic [][]Pipe, sx, sy int) int {
	fmt.Printf("Start x: %v y: %v\n", sx, sy)
	currentPipe, direction := getStart(schematic, sx, sy)
	switch direction {
	case Down:
		sy--
	case Up:
		sy++
	case Right:
		sx--
	case Left:
		sx++
	}
	curX, curY := sx, sy
	fmt.Printf("pipe: %v coord x: %v, y: %v direction: %v\n", currentPipe.s, curX, curY, directionToString[direction])
	result := 1
	for currentPipe.s != "S" {
		x, y, d := currentPipe.next(direction)
		curX += x
		curY += y
		direction = d
		currentPipe = schematic[curY][curX]
		fmt.Printf("pipe: %v coord x: %v, y: %v direction: %v\n", currentPipe.s, curX, curY, directionToString[direction])
		result++
	}

	return result / 2
}

func (p *Pipe) next(d Direction) (x, y int, direction Direction) {
	//d is direction you are coming from
	switch p.s {
	case "-":
		if d == Right {
			x = -1
			direction = Right
		} else {
			x = 1
			direction = Left
		}
	case "|":
		if d == Down {
			y = -1
			direction = Down
		} else {
			y = 1
			direction = Up
		}
	case "7":
		if d == Down {
			x = -1
			direction = Right
		} else {
			y = 1
			direction = Up
		}
	case "J":
		if d == Up {
			x = -1
			direction = Right
		} else {
			y = -1
			direction = Down
		}
	case "L":
		if d == Up {
			x = 1
			direction = Left
		} else {
			y = -1
			direction = Down
		}
	case "F":
		if d == Down {
			x = 1
			direction = Left
		} else {
			y = 1
			direction = Up
		}

	}
	return x, y, direction
}

func getStart(schematic [][]Pipe, sx, sy int) (Pipe, Direction) {
	right := schematic[sy][sx+1]
	down := schematic[sy+1][sx]
	up := schematic[sy-1][sx]
	if up.a == Down || up.b == Down {
		return up, Down
	}
	if right.a == Left || right.b == Left {
		return right, Left
	} else if down.a == Up || down.b == Up {
		return down, Up
	}
	left := schematic[sy][sx-1]
	return left, Right
}

type Pipe struct {
	a, b Direction
	s    string
}

var pipemap = map[rune]Pipe{
	'-': {a: Left, b: Right, s: "-"},
	'|': {a: Up, b: Down, s: "|"},
	'7': {a: Down, b: Left, s: "7"},
	'J': {a: Up, b: Left, s: "J"},
	'L': {a: Up, b: Right, s: "L"},
	'F': {a: Down, b: Right, s: "F"},
	'S': {s: "S"},
}

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

var directionToString = []string{"Left", "Right", "Up", "Down"}
