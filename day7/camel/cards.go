package camel

import (
	"slices"
	"strconv"
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

var FaceCardValues = map[rune]int{
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

const (
	FiveOfAKind  string = "FiveOfAKind"
	FourOfAKind  string = "FourOfAKind"
	FullHouse    string = "FullHouse"
	ThreeOfAKind string = "ThreeOfAKind"
	TwoPairs     string = "TwoPairs"
	OnePair      string = "OnePair"
	HighCard     string = "HighCard"
	Nothing      string = "Nothing"
)

var HandValues = map[string]int{
	FiveOfAKind:  6,
	FourOfAKind:  5,
	FullHouse:    4,
	ThreeOfAKind: 3,
	TwoPairs:     2,
	OnePair:      1,
	HighCard:     0,
}

type Hand struct {
	Cards      string
	Bid        int
	HandValue  int
	CardValues []int
}

func CountWinnings(input []string, useJokers bool) int {
	evaluate := evaluateHand
	if useJokers {
		FaceCardValues['J'] = 1
		evaluate = evaluateHandWithJokers
	}

	hands := []Hand{}
	for _, line := range input {
		splits := strings.Split(line, " ")
		cards := splits[0]
		bid, _ := strconv.Atoi(splits[1])
		handValue, cardValues := evaluate(cards)
		hands = append(hands, Hand{cards, bid, handValue, cardValues})
	}

	slices.SortFunc(hands, compareHands)

	total := 0
	for i, hand := range hands {
		rank := i + 1
		total += rank * hand.Bid
	}

	return total
}

func countsAndValues(cards string) (map[rune]int, []int) {
	counts := map[rune]int{}
	values := []int{}
	for _, card := range cards {
		count, ok := counts[card]
		if !ok {
			counts[card] = 1
		} else {
			counts[card] = count + 1
		}

		value, ok := kit.RuneToInt(card)
		if !ok {
			value = FaceCardValues[card]
		}
		values = append(values, value)
	}

	return counts, values
}

func evaluateHand(cards string) (int, []int) {
	counts, values := countsAndValues(cards)
	kinds, pairs := 0, 0
	for _, count := range counts {
		if kinds < count {
			kinds = count
		}

		if count == 2 {
			pairs++
		}
	}

	eval := 0
	switch {
	case kinds == 5:
		eval = HandValues[FiveOfAKind]
	case kinds == 4:
		eval = HandValues[FourOfAKind]
	case kinds == 3 && pairs == 1:
		eval = HandValues[FullHouse]
	case kinds == 3:
		eval = HandValues[ThreeOfAKind]
	case pairs == 2:
		eval = HandValues[TwoPairs]
	case pairs == 1:
		eval = HandValues[OnePair]
	default:
		eval = HandValues[HighCard]
	}

	return eval, values
}

func evaluateHandWithJokers(cards string) (int, []int) {
	counts, values := countsAndValues(cards)
	kinds, pairs, jokers := 0, 0, 0
	for key, count := range counts {
		if key == 'J' {
			jokers = count
			continue
		}

		if kinds < count {
			kinds = count
		}

		if count == 2 {
			pairs++
		}
	}

	eval := 0
	switch {
	case kinds+jokers == 5:
		eval = HandValues[FiveOfAKind]
	case kinds+jokers == 4:
		eval = HandValues[FourOfAKind]
	case kinds == 3 && pairs == 1, kinds == 3 && jokers == 1, pairs == 2 && jokers == 1, pairs == 1 && jokers >= 2:
		eval = HandValues[FullHouse]
	case kinds+jokers == 3, kinds == 2 && jokers >= 1:
		eval = HandValues[ThreeOfAKind]
	case pairs == 2, pairs == 1 && jokers == 1:
		eval = HandValues[TwoPairs]
	case pairs == 1, pairs == 0 && jokers == 1:
		eval = HandValues[OnePair]
	default:
		eval = HandValues[HighCard]
	}

	return eval, values
}

func compareHands(a, b Hand) int {
	if a.HandValue != b.HandValue {
		return a.HandValue - b.HandValue
	}

	for i := 0; i < len(a.CardValues); i++ {
		avalue := a.CardValues[i]
		bvalue := b.CardValues[i]
		if avalue != bvalue {
			return avalue - bvalue
		}
	}

	return 0
}
