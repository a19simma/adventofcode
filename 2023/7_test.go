package adventofcode

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var valueMap = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}
var valueMap2 = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 1,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func TestSimpleCamelCards(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	game := ParseGame(input)
	result := CalculateWinnings(&game)
	assert.Equal(t, 6440, result)
}
func TestCamelCards(t *testing.T) {
	input, _ := os.ReadFile("7.txt")
	game := ParseGame(string(input))
	result := CalculateWinnings(&game)
	assert.Equal(t, 248422077, result)
}
func TestSimpleCamelCardsJoker(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	game := ParseGame2(input)
	result := CalculateWinnings(&game)
	assert.Equal(t, 5905, result)
}
func TestCamelCardsJoker(t *testing.T) {
	input, _ := os.ReadFile("7.txt")
	game := ParseGame2(string(input))
	result := CalculateWinnings(&game)
	assert.Equal(t, 249817836, result)
}

func TestInsertIntoHand2(t *testing.T) {
	input := `J4334 765`
	game := ParseGame2(input)
	CalculateWinnings(&game)
	assert.Equal(t, 1, len(game.house))
}

func CalculateWinnings(game *Game) int {
	result := 0
	i := 1
	for _, v := range game.high {
		result += i * v.bid
		fmt.Printf("cards: %v strength: %v\n", string(v.cards), 0)
		i++
	}
	for _, v := range game.onePair {
		result += i * v.bid
		fmt.Printf("cards: %v strength: %v\n", string(v.cards), 1)
		i++
	}
	for _, v := range game.twoPair {
		result += i * v.bid
		fmt.Printf("cards: %v strength: %v\n", string(v.cards), 2)
		i++
	}
	for _, v := range game.three {
		result += i * v.bid
		fmt.Printf("cards: %v strength: %v\n", string(v.cards), 3)
		i++
	}
	for _, v := range game.house {
		result += i * v.bid
		fmt.Printf("cards: %v strength: %v\n", string(v.cards), 4)
		i++
	}
	for _, v := range game.four {
		result += i * v.bid
		fmt.Printf("cards: %v strength: %v\n", string(v.cards), 5)
		i++
	}
	for _, v := range game.five {
		result += i * v.bid
		fmt.Printf("cards: %v strength: %v\n", string(v.cards), 6)
		i++
	}
	return result
}

func ParseGame(input string) Game {
	game := Game{}
	for _, h := range strings.Split(input, "\n") {
		values := strings.Split(h, " ")
		if len(values) < 2 {
			continue
		}
		c := values[0]
		b, _ := strconv.Atoi(values[1])
		runes := []rune(c)
		hand := Hand{cards: runes, bid: b}
		game.InsertHand(hand)
	}
	return game
}
func ParseGame2(input string) Game {
	game := Game{}
	for _, h := range strings.Split(input, "\n") {
		values := strings.Split(h, " ")
		if len(values) < 2 {
			continue
		}
		c := values[0]
		b, _ := strconv.Atoi(values[1])
		runes := []rune(c)
		hand := Hand{cards: runes, bid: b}
		game.InsertHand2(hand)
	}
	return game
}

func (g *Game) InsertHand(h Hand) {
	for _, c := range h.cards {
		switch count(c, h.cards) {
		case 5:
			g.five = insertIntoHandArray(h, g.five)
			return
		case 4:
			g.four = insertIntoHandArray(h, g.four)
			return
		case 3:
			var hasPair bool
			for _, card := range h.cards {
				if count(card, h.cards) == 2 && c != card {
					hasPair = true
					break
				}
			}
			if hasPair {
				g.house = insertIntoHandArray(h, g.house)
			} else {
				g.three = insertIntoHandArray(h, g.three)
			}
			return
		case 2:
			var hasPair, hasThree bool
			for _, card := range h.cards {
				if count(card, h.cards) == 3 && c != card {
					hasThree = true
					break
				}
				if count(card, h.cards) == 2 && c != card {
					hasPair = true
					break
				}
			}
			if hasThree {
				g.house = insertIntoHandArray(h, g.house)
			} else if hasPair {
				g.twoPair = insertIntoHandArray(h, g.twoPair)
			} else {
				g.onePair = insertIntoHandArray(h, g.onePair)
			}
			return
		}
	}
	g.high = insertIntoHandArray(h, g.high)
}
func (g *Game) InsertHand2(h Hand) {
	cardset := make(map[rune]int)
	for _, c := range h.cards {
		cc := count(c, h.cards)
		cardset[c] = cc
	}

	var maxRune rune
	curMaxVal := 0
	for k, v := range cardset {
		if v > curMaxVal && k != 'J' {
			curMaxVal = v
			maxRune = k
		}
	}

	cardset[maxRune] += cardset['J']
	delete(cardset, 'J')

	switch len(cardset) {
	case 1:
		g.five = insertIntoHandArray2(h, g.five)
		return
	case 2:
		var numFour, numThree int
		for _, v := range cardset {
			if v == 4 {
				numFour++
			}
			if v == 3 {
				numThree++
			}
		}
		if numFour == 1 {
			g.four = insertIntoHandArray2(h, g.four)
		} else if numThree == 1 {
			g.house = insertIntoHandArray2(h, g.house)
		}
		return
	case 3, 4:
		var numPair, numThree int
		for _, v := range cardset {
			if v == 2 {
				numPair++
			}
			if v == 3 {
				numThree++
			}
		}
		if numPair == 1 && numThree == 1 {
			g.house = insertIntoHandArray2(h, g.house)
		} else if numPair == 2 {
			g.twoPair = insertIntoHandArray2(h, g.twoPair)
		} else if numPair == 1 {
			g.onePair = insertIntoHandArray2(h, g.onePair)
		} else if numThree == 1 {
			g.three = insertIntoHandArray2(h, g.three)
		}
		return

	case 5:
		g.high = insertIntoHandArray2(h, g.high)
		return
	}
}

func insertIntoHandArray(hand Hand, handArray []Hand) []Hand {
	for i, h := range handArray {
		for j, c := range h.cards {
			if valueMap[hand.cards[j]] == valueMap[c] {
				continue
			}
			if valueMap[hand.cards[j]] < valueMap[c] {
				newArray := []Hand{}
				newArray = append(newArray, handArray[:i]...)
				newArray = append(newArray, hand)
				newArray = append(newArray, handArray[i:]...)
				return newArray
			} else {
				break
			}
		}
	}
	return append(handArray, hand)
}
func insertIntoHandArray2(hand Hand, handArray []Hand) []Hand {
	for i, h := range handArray {
		for j, c := range h.cards {
			if valueMap2[hand.cards[j]] == valueMap2[c] {
				continue
			}
			if valueMap2[hand.cards[j]] < valueMap2[c] {
				newArray := []Hand{}
				newArray = append(newArray, handArray[:i]...)
				newArray = append(newArray, hand)
				newArray = append(newArray, handArray[i:]...)
				return newArray
			} else {
				break
			}
		}
	}
	return append(handArray, hand)
}

func count(rune rune, runes []rune) int {
	result := 0
	for _, r := range runes {
		if r == rune {
			result++
		}
	}
	return result
}
func count2(inRune rune, runes []rune) int {
	tmpRunes := []rune{}
	tmpRunes = append(tmpRunes, runes...)

	cardset := make(map[rune]int)
	for _, c := range tmpRunes {
		cc := count(c, tmpRunes)
		cardset[c] = cc
	}

	var maxRune rune
	curMaxVal := 0
	for k, v := range cardset {
		if v > curMaxVal && k != 'J' {
			curMaxVal = v
			maxRune = k
		}
	}
	for i, r := range tmpRunes {
		if r == 'J' {
			tmpRunes[i] = maxRune
		}
	}
	result := 0
	for _, r := range tmpRunes {
		if r == inRune {
			result++
		}
	}
	return result
}

type Hand struct {
	cards []rune
	bid   int
}

type Game struct {
	five, four, house, three, twoPair, onePair, high []Hand
}
