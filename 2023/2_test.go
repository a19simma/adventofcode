package adventofcode2023

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Draw struct {
	red   int
	blue  int
	green int
}

func TestSimpleCubeGame(t *testing.T) {
	data := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	games := GetGames(data)
	result := CubeGame(Draw{red: 12, green: 13, blue: 14}, games)
	assert.Equal(t, 8, result)
}

func TestCubeGame(t *testing.T) {
	data, _ := os.ReadFile("2.txt")
	games := GetGames(string(data))
	result := CubeGame(Draw{red: 12, green: 13, blue: 14}, games)
	assert.Equal(t, 2156, result)
}

func TestSimpleCubeMinPower(t *testing.T) {
	data := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	games := GetGames(data)
	result := MinCubePower(games)
	assert.Equal(t, 2286, result)
}

func TestCubeMinPower(t *testing.T) {
	data, _ := os.ReadFile("2.txt")
	games := GetGames(string(data))
	result := MinCubePower(games)
	assert.Equal(t, 66909, result)
}

func MinCubePower(games [][]Draw) int {
	result := 0
	for _, draws := range games {
		minDraw := Draw{}
		for _, draw := range draws {
			minDraw.ReplaceIfGreater(draw)
		}
		result += minDraw.red * minDraw.green * minDraw.blue
	}
	return result
}

func (d *Draw) ReplaceIfGreater(in Draw) {
	if in.red > d.red {
		d.red = in.red
	}
	if in.green > d.green {
		d.green = in.green
	}
	if in.blue > d.blue {
		d.blue = in.blue
	}
}

func GetGames(data string) [][]Draw {
	lines := strings.Split(string(data), "\n")
	games := [][]Draw{}
	for _, l := range lines {
		game := []Draw{}
		gameLine := strings.Split(l, ":")
		if len(gameLine) < 2 {
			continue
		}
		gameDraws := strings.Split(gameLine[1], ";")
		for _, d := range gameDraws {
			balls := strings.Split(d, ",")
			draw := Draw{}
			for _, b := range balls {
				trim := strings.Trim(b, " ")
				split := strings.Split(trim, " ")
				scoreString := split[0]
				score, _ := strconv.Atoi(scoreString)
				color := split[1]
				switch color {
				case "blue":
					draw.blue = score
				case "red":
					draw.red = score
				case "green":
					draw.green = score
				}
			}
			game = append(game, draw)
		}
		games = append(games, game)
	}
	return games
}

func CubeGame(conf Draw, games [][]Draw) int {
	result := 0
	for i, v := range games {
		possible := true
		for _, d := range v {
			if d.blue > conf.blue || d.green > conf.green || d.red > conf.red {
				possible = false
				break
			}

		}
		if possible {
			result += i + 1
		}
	}
	return result
}
