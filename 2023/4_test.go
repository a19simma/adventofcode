package adventofcode

import (
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleScratchCard(t *testing.T) {
	data := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	result := GetScratchCardScore(data)
	assert.Equal(t, 13, result)
}

func TestScratchCard(t *testing.T) {
	data, _ := os.ReadFile("4.txt")
	result := GetScratchCardScore(string(data))
	assert.Equal(t, 25004, result)
}

func TestSimpleRecursiveScratchCard(t *testing.T) {
	data := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	cardConfig := GetCardConfig(data)
	result := GetScratchCardScoreRecursive(cardConfig, &cardConfig, 0)
	assert.Equal(t, 30, result)
}

func TestRecursiveScratchCard(t *testing.T) {
	data, _ := os.ReadFile("4.txt")
	cardConfig := GetCardConfig(string(data))
	result := GetScratchCardScoreRecursive(cardConfig, &cardConfig, 0)
	assert.Equal(t, 14427616, result)
}

type CardConfigs struct {
	winnings, cards []int
}

func GetScratchCardScoreRecursive(cardConfig []CardConfigs, allCards *[]CardConfigs, index int) int {
	result := 0
	configLength := len(*allCards)
	for i, cc := range cardConfig {
		score := 0
		for _, card := range cc.cards {
			if slices.Contains(cc.winnings, card) {
				score += 1
			}
		}
		result += 1
		if score == 0 {
			continue
		}
		score += i + index + 1
		if score > configLength {
			score = configLength
		}
		tmp := (*allCards)[index+i+1 : score]
		result += GetScratchCardScoreRecursive(tmp, allCards, i+index+1)
	}
	return result
}

func GetCardConfig(data string) []CardConfigs {
	cardConfigs := []CardConfigs{}
	for _, v := range strings.Split(data, "\n") {
		values := strings.Split(v, ":")
		if len(values) < 2 {
			continue
		}
		cards := strings.Split(values[1], "|")
		winning := []int{}
		for _, w := range strings.Fields(cards[0]) {
			i, _ := strconv.Atoi(w)
			winning = append(winning, i)
		}
		selectedCards := []int{}
		if len(cards) > 1 {
			for _, c := range strings.Fields(cards[1]) {
				i, _ := strconv.Atoi(c)
				selectedCards = append(selectedCards, i)

			}
		}
		cardConfigs = append(cardConfigs, CardConfigs{winnings: winning, cards: selectedCards})
	}
	return cardConfigs
}
func GetScratchCardScore(data string) int {
	result := 0
	cardConfigs := GetCardConfig(data)
	for _, cc := range cardConfigs {
		score := 0
		for _, card := range cc.cards {
			if slices.Contains(cc.winnings, card) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		result += score
	}
	return result
}
