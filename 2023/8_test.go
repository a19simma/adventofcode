package adventofcode

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	left, right, label string
	start, end         bool
}

type direction int

const (
	left direction = iota
	right
)

func TestSimpleMap(t *testing.T) {
	input := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
	desertMap, instructionSet := ParseMap(input)
	result := FindEnd(desertMap, instructionSet)
	assert.Equal(t, 6, result)
}
func TestMap(t *testing.T) {
	input, _ := os.ReadFile("8.txt")
	desertMap, instructionSet := ParseMap(string(input))
	result := FindEnd(desertMap, instructionSet)
	assert.Equal(t, 11911, result)
}

func TestSimleGhostMap(t *testing.T) {
	input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
	desertMap, instructionSet := ParseMap(string(input))
	result := FindEnd2(desertMap, instructionSet)
	assert.Equal(t, 6, result)

}
func TestGhostMap(t *testing.T) {
	input, _ := os.ReadFile("8.txt")
	desertMap, instructionSet := ParseMap(string(input))
	result := FindEnd2(desertMap, instructionSet)
	assert.Equal(t, 6, result)
}
func FindEnd2(desertMap map[string]Node, instructionSet []direction) int {
	startNodes := []Node{}
	for _, n := range desertMap {
		if n.start {
			startNodes = append(startNodes, n)
		}
	}
	steps := make([]int, len(startNodes))
	for i := 0; i < len(startNodes); i++ {
		for j := 0; !startNodes[i].end; j++ {
			instruction := instructionSet[j%len(instructionSet)]
			if instruction == left {
				startNodes[i] = desertMap[startNodes[i].left]
			} else {
				startNodes[i] = desertMap[startNodes[i].right]
			}
			steps[i] = j
		}
	}
	if len(steps) == 2 {
		return LCM(steps[0], steps[1])
	}

	return LCM(steps[0], steps[1], steps[:2]...)
}

func FindEnd(desertMap map[string]Node, instructionSet []direction) int {
	curNode := desertMap["AAA"]
	result := 0
	i := 0
	for curNode.label != "ZZZ" {
		instruction := instructionSet[i%len(instructionSet)]
		if instruction == left {
			curNode = desertMap[curNode.left]
		} else {
			curNode = desertMap[curNode.right]
		}
		i++
		result++
	}

	return result
}

func ParseMap(input string) (map[string]Node, []direction) {
	sections := strings.Split(input, "\n\n")
	instructions := []direction{}
	for _, r := range sections[0] {
		switch r {
		case 'R':
			instructions = append(instructions, right)
		case 'L':
			instructions = append(instructions, left)
		}
	}
	nodes := make(map[string]Node)
	for _, n := range strings.Split(sections[1], "\n") {
		label := string([]rune(n)[:3])
		left := string([]rune(n)[7:10])
		right := string([]rune(n)[12:15])
		node := Node{label: label, left: left, right: right}

		if string(label[2]) == "Z" {
			node.end = true
		}
		if string(label[2]) == "A" {
			node.start = true
		}
		nodes[label] = node
	}
	for i, n := range nodes {
		fmt.Printf("nodes: %v index: %v\n", n, i)
	}
	fmt.Printf("instructions: %v\n", instructions)
	return nodes, instructions
}
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int, ints ...int) int {
	result := a * b / GCD(a, b)

	for _, i := range ints {
		result = LCM(result, i)
	}

	return result
}
